package database

import (
	"consumer-service/internal/entity"
	"database/sql"
	"fmt"
)

type AccountDb struct {
	Db *sql.DB
}

func NewAccountDb(db *sql.DB) *AccountDb {
	return &AccountDb{
		Db: db,
	}
}

func (c *AccountDb) Create(account *entity.Account) error {
	smtp, err := c.Db.Prepare("insert into accounts (AccountID, AccountBalance) values (?,?)")
	if err != nil {
		return err
	}
	defer smtp.Close()
	_, err = smtp.Exec(account.AccountId, account.AccountBalance)
	if err != nil {
		return err
	}
	return nil
}

func (c *AccountDb) FindById(id string) (*entity.Account, error) {
	var account = &entity.Account{}
	smtp, err := c.Db.Prepare("select AccountID,AccountBalance from accounts where AccountID = ?")
	fmt.Println("select", err)
	if err != nil {
		return nil, err
	}
	defer smtp.Close()
	row := smtp.QueryRow(id)
	if err = row.Scan(&account.AccountId, &account.AccountBalance); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *AccountDb) Update(account *entity.Account) error {
	smtp, err := c.Db.Prepare("update accounts set AccountBalance = ? where AccountID = ?")
	if err != nil {
		return err
	}
	defer smtp.Close()
	_, err = smtp.Exec(account.AccountBalance, account.AccountId)
	if err != nil {
		return err
	}
	return nil
}

// INSERT INTO cliente (id, nome) VALUES (10, 'nome do cliente') ON DUPLICATE KEY UPDATE nome = 'nome do cliente';
func (c *AccountDb) Update2(account *entity.Account) error {
	smtp, err := c.Db.Prepare("REPLACE INTO accounts  (AccountID, AccountBalance)  values (?,?) ")
	if err != nil {
		return err
	}
	defer smtp.Close()
	_, err = smtp.Exec(account.AccountId, account.AccountBalance)
	if err != nil {
		return err
	}
	return nil
}
