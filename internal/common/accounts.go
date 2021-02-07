package common

import "strings"

type Email string
type Username string

type Account struct {
	Email            Email
	ExternalUsername Username
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
		emails = append(emails, string(acc.Email))
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
				ExternalUsername: Username(strings.TrimSpace(accFields[1])),
			},
		)
	}

	return accounts
}
