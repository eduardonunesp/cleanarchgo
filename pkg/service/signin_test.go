package service

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignin(t *testing.T) {
	suite.Run(t, new(testSigninSuite))
}

type testSigninSuite struct {
	suite.Suite
	// Mocks goes here
	useCase *Signin
}

func (s *testSigninSuite) SetupTest() {
	// Mocks goes here
}

func (s *testSigninSuite) TestSigninSuccess() {
	// Test goes here
}
