package domain

import "github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"

func BuildAccount(options ...accountOption) (*Account, error) {
	newAcc := Account{}
	for _, opt := range options {
		if err := opt(&newAcc); err != nil {
			return nil, RaiseDomainError(err)
		}
	}
	return &newAcc, nil
}

func AccountWithNewID() accountOption {
	return func(opt *Account) error {
		opt.id = valueobject.MustUUID()
		return nil
	}
}

func AccountWithID(id string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.id, err = valueobject.UUIDFromString(id)
		return err
	}
}

func AccountWithName(name string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.name, err = valueobject.NameFromString(name)
		return err
	}
}

func AccountWithEmail(email string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.email, err = valueobject.EmailFromString(email)
		return err
	}
}

func AccountWithCpf(cpf string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.cpf, err = valueobject.CpfFromString(cpf)
		return err
	}
}

func AccountWithCarPlate(carPlate string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.carPlate, err = valueobject.CarPlateFromString(carPlate)
		return err
	}
}

func AccountWithHash(hash string) accountOption {
	return func(opt *Account) error {
		opt.hash = valueobject.LoadHashFromString(hash)
		return nil
	}
}

func AccountWithPassword(password string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.hash, err = valueobject.BuildHashFromString(password, nil)
		return err
	}
}

func AccountWithAccountType(accountType string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.accountType, err = valueobject.AccountTypeFromString(accountType)
		return err
	}
}

func AccountWithConfirmedAt(date int64) accountOption {
	return func(opt *Account) error {
		var err error
		opt.confirmedAt, err = valueobject.DateFromUnix(date)
		return err
	}
}

func AccountWithUpdatedAt(date int64) accountOption {
	return func(opt *Account) error {
		var err error
		opt.updatedAt, err = valueobject.DateFromUnix(date)
		return err
	}
}

func AccountWithCreatedAt(date int64) accountOption {
	return func(opt *Account) error {
		var err error
		opt.createdAt, err = valueobject.DateFromUnix(date)
		return err
	}
}
