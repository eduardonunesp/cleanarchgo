package service

type SigninParams struct {
	// Params goes here
}

type SigninResult struct {
	// Result goes here
}

type Signin struct {
	// Dependencies goes here
}

func NewSignin() *Signin {
	return &Signin{}
}
func (s Signin) Execute(input *SigninParams) (*SigninResult, error) {
	return nil, nil
}
