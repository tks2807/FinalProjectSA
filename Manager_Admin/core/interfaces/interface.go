package interfaces

import "Manager_Admin/core"

type IAccountRepository interface {
	GetAll() []*core.Account
	CreateAccount(account core.Account) bool
}
