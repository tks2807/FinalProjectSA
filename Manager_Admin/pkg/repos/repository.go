package repos

import (
	"Manager_Admin/core"
	"Manager_Admin/core/interfaces"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AccountRepository struct {
	pool pgxpool.Pool
}

func newAccountRepository(conn *pgxpool.Pool) interfaces.IAccountRepository {
	return &AccountRepository{*conn}
}

func (r AccountRepository) GetAll() []*core.Account{
	stmt := "SELECT * FROM account"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		return nil
	}
	defer rows.Close()
	accs := []*core.Account{}
	for rows.Next() {
		a := &core.Account{}
		err = rows.Scan(&a.Id, &a.Name, &a.Surname, &a.Email)
		if err != nil {
			return nil
		}
		accs = append(accs, a)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return accs
}

func (c *AccountRepository) CreateAccount(acc core.Account) bool {
	sql := "INSERT INTO account(id, name, surname, email) " +
		"VALUES($1, $2, $3, $4)";
	row := c.pool.QueryRow(context.Background(),
		sql, acc.Id, acc.Name, acc.Surname, acc.Email)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}

