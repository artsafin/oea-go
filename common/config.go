package common

import "strings"

type Config struct {
	BaseUri            string `mapstructure:"base_uri"`
	ApiTokenOf         string `mapstructure:"api_token_of"`
	ApiTokenEm         string `mapstructure:"api_token_em"`
	DocIdOf            string `mapstructure:"doc_id_of"`
	DocIdEm            string `mapstructure:"doc_id_em"`
	AuthKey            string `mapstructure:"auth_key"`
	AuthAllowedDomains string `mapstructure:"auth_allowed_domains"`
	AuthAllowedEmails  string `mapstructure:"auth_allowed_emails"`
	SmtpHost           string `mapstructure:"smtp_host"`
	SmtpUser           string `mapstructure:"smtp_user"`
	SmtpPass           string `mapstructure:"smtp_pass"`
	SmtpPort           int
}

func NewConfig() Config {
	return Config{SmtpPort: 25}
}

func (c Config) MustValidate() {
	if c == (Config{}) {
		panic("config validation: not all config parameters are set")
	}
}

func extractEmailDomain(email string) string {
	if emailDomainIdx := strings.Index(email, "@"); emailDomainIdx >= 0 {
		return email[emailDomainIdx+1:]
	}
	return email
}

func (c Config) isAuthAllowedForDomain(email string) bool {
	domains := strings.Split(c.AuthAllowedDomains, ",")

	emailDomain := extractEmailDomain(email)

	for _, domain := range domains {
		if strings.TrimSpace(domain) == emailDomain {
			return true
		}
	}
	return false
}

func (c Config) isAuthAllowedForEmail(email string) bool {
	emails := strings.Split(c.AuthAllowedEmails, ",")

	for _, allowedEmail := range emails {
		if strings.TrimSpace(allowedEmail) == email {
			return true
		}
	}
	return false
}

func (c Config) IsAuthAllowed(email string) bool {
	return c.isAuthAllowedForEmail(email) || c.isAuthAllowedForDomain(email)
}
