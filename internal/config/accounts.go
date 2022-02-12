package config

import (
	"fmt"
	"strings"
)

type Email string
type Username string

type Account struct {
	Email            Email
	ExternalUsername Username
}

func (a Account) String() string {
	return fmt.Sprintf("<%v:%v>", a.Email, a.ExternalUsername)
}

type Accounts []Account

func (a Accounts) HasEmail(email string) bool {
	return a.Get(email) != nil
}

func (a Accounts) Get(email string) *Account {
	for _, acc := range a {
		if acc.Email == Email(email) {
			return &acc
		}
	}
	return nil
}

func (a Accounts) String() string {
	emails := make([]string, 0, len(a))
	for _, acc := range a {
		emails = append(emails, acc.String())
	}

	return "Accounts{" + strings.Join(emails, ", ") + "}"
}

/*
	Input: "test@example.com:username, test2@example.com:foo"
	Output: Accounts{
		Account{Email: "test@example.com", ExternalUsername: "username"},
		Account{Email: "test2@example.com", ExternalUsername: "foo"},
	}
*/
func newAccountsFromConfig(authAccounts string) Accounts {
	accountDefs := strings.Split(authAccounts, ",")
	accounts := make(Accounts, 0, len(accountDefs))
	for _, accDef := range accountDefs {
		accFields := strings.Split(strings.TrimSpace(accDef), ":")

		if len(accFields) != 2 {
			continue
		}

		accounts = append(
			accounts,
			Account{
				Email:            Email(strings.TrimSpace(accFields[0])),
				ExternalUsername: NewUsernameFromString(accFields[1]),
			},
		)
	}

	return accounts
}

func NewUsernameFromString(strUsername string) Username {
	return Username(
		strings.ToLower(
			strings.TrimLeft(
				strings.TrimSpace(strUsername),
				"@",
			)),
	)
}
