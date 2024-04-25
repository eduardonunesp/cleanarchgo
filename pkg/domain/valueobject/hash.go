package valueobject

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	ErrHashRawValueEmpty   = errors.New("raw value cannot be empty")
)

// DefaultParams provides some sane default parameters for hashing passwords.
//
// Follows recommendations given by the Argon2 RFC:
// "The Argon2id variant with t=1 and maximum available memory is RECOMMENDED as a
// default setting for all environments. This setting is secure against side-channel
// attacks and maximizes adversarial costs on dedicated bruteforce hardware.""
//
// The default parameters should generally be used for development/testing purposes
// only. Custom parameters should be set for production applications depending on
var DefaultParams = &HashParams{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

type HashParams struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32
	// The number of iterations over the memory.
	Iterations uint32
	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8
	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32
	// Length of the generated key. 16 bytes or more is recommended.
	KeyLength uint32
}

type Hash struct {
	// The encoded hash value
	value  string
	params *HashParams
}

// LoadHashFromString loads hash from the given encoded hash
func LoadHashFromString(encodedValue string) Hash {
	return Hash{encodedValue, DefaultParams}
}

// BuildHashFromString generates a new hash from the given password string using
// the provided parameters. If no parameters are provided, the DefaultParams
func BuildHashFromString(rawValue string, params *HashParams) (Hash, error) {
	if rawValue == "" {
		return Hash{}, ErrHashRawValueEmpty
	}
	if params == nil {
		params = DefaultParams
	}
	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	encodedHash, err := generateFromPassword(rawValue, params)
	if err != nil {
		log.Fatal(err)
	}
	return Hash{encodedHash, params}, nil
}

func (h Hash) String() string {
	return h.value
}

func (h Hash) Compare(password string) (bool, error) {
	return comparePasswordAndHash(password, h.value)
}

func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func generateFromPassword(password string, p *HashParams) (string, error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}
	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	// Return a string using the standard encoded hash representation.
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash), nil
}

func decodeHash(encodedHash string) (p *HashParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}
	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}
	p = &HashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))
	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))
	return p, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
