package repository

import (
	"context"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type AccountRepositoryPG struct {
	conn *pgx.Conn
}

func NewAccountRepositoryPG(connStr string) *AccountRepositoryPG {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return &AccountRepositoryPG{conn}
}

func (a AccountRepositoryPG) HasAccountByEmail(email string) (bool, error) {
	queries := db.New(a.conn)
	result, err := queries.HasAccountByEmail(context.Background(), email)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	return result, nil
}

func (a AccountRepositoryPG) GetAccountByID(id string) (*domain.Account, error) {
	queries := db.New(a.conn)
	uuid, err := mapStringToPgTypeUUID(id)
	if err != nil {
		return nil, RaiseRepositoryError(err)
	}
	account, err := queries.GetAccount(context.Background(), uuid)
	if err != nil {
		return nil, RaiseRepositoryError(err)
	}
	return mapDBAccountToDomainAccount(&account)
}

func (a AccountRepositoryPG) SaveAccount(account *domain.Account) error {
	queries := db.New(a.conn)
	saveAccountParams, err := mapDomainAccountToSaveAccountParams(account)
	if err != nil {
		return RaiseRepositoryError(err)
	}
	err = queries.SaveAccount(context.Background(), *saveAccountParams)
	if err != nil {
		return RaiseRepositoryError(err)
	}
	return nil
}

func (a AccountRepositoryPG) IsDriverFreeByDriverID(driverID string) (bool, error) {
	queries := db.New(a.conn)
	uuid, err := mapStringToPgTypeUUID(driverID)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	result, err := queries.IsDriverFree(context.Background(), uuid)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	return result, nil
}
