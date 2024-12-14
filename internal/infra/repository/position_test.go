//go:build db_integration

package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPositionRepositoryPG(t *testing.T) {
	suite.Run(t, new(testPositionRepoPGSuite))
}

type testPositionRepoPGSuite struct {
	suite.Suite
	repoDB *PositionRepositoryPG
}

func (s *testPositionRepoPGSuite) SetupTest() {
	s.repoDB = NewPositionRepositoryPG("postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
}
