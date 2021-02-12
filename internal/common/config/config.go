package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io"
	"strings"
)

func boolStr(val bool, trueVal, falseVal string) string {
	if val {
		return trueVal
	}
	return falseVal
}

type Config struct {
	AppVersion string   `ignored:"true"`
	Accounts   Accounts `ignored:"true"`
	SecretKey  []byte   `ignored:"true"`

	BaseUri      string `envconfig:"base_uri" required:"true"`
	ApiTokenOf   string `envconfig:"api_token_of" required:"true"`
	ApiTokenEm   string `envconfig:"api_token_em" required:"true"`
	DocIdOf      string `envconfig:"doc_id_of" required:"true"`
	DocIdEm      string `envconfig:"doc_id_em" required:"true"`
	RawSecretKey string `envconfig:"secret_key" required:"true"`
	RawAccounts  string `envconfig:"auth_accounts" required:"true"`
	SmtpHost     string `envconfig:"smtp_host" required:"true"`
	SmtpUser     string `envconfig:"smtp_user" required:"true"`
	SmtpPass     string `envconfig:"smtp_pass" required:"true"`
	SmtpPort     uint   `envconfig:"smtp_port" required:"true"`
	SecurePort   uint   `envconfig:"secure_port"`
	InsecurePort uint   `envconfig:"insecure_port"`
	TlsCert      string `envconfig:"tls_cert"`
	TlsKey       string `envconfig:"tls_key"`
	UseAuth      bool   `envconfig:"use_auth" required:"true"`
	BotToken     string `envconfig:"bot_token"`
	FilesDir     string `envconfig:"files"`
}

func NewDefaultConfig(appVersion string) Config {
	return Config{
		BaseUri:      "https://coda.io/apis/v1beta1",
		SecurePort:   8443,
		InsecurePort: 8080,
		SmtpPort:     25,
		UseAuth:      true,
		AppVersion:   appVersion,
	}
}

func (c *Config) DumpNonSecretParameters(wr io.Writer) {
	lines := []string{
		"\033[31mUse Auth:\033[0m %v",
		"\033[31mTLS:\033[0m %v",
		"\033[33mDocument IDs:\033[0m of=%v; em=%v",
		"\033[33mUser accounts:\033[0m %v",
		"\033[33mSmtp:\033[0m %v (%v)",
		"",
	}
	params := []interface{}{
		c.UseAuth,
		boolStr(c.TlsCert != "", "configured", "disabled"),
		c.DocIdOf,
		c.DocIdEm,
		c.Accounts,
		boolStr(c.SmtpHost != "", fmt.Sprintf("%v:%v", c.SmtpHost, c.SmtpPort), "email sending disabled"),
		boolStr(c.SmtpHost != "" && c.SmtpPass != "", "auth enabled", "auth disabled"),
	}
	_, _ = wr.Write([]byte(fmt.Sprintf(strings.Join(lines, "\n"), params...)))
}

func (c *Config) IsTLS() bool {
	return c.TlsCert != ""
}

func (c *Config) LoadFromEnvAndValidate() error {
	err := envconfig.Process("oea", c)

	if err != nil {
		return fmt.Errorf("config: error loading config: %v", err)
	}

	c.Accounts = newAccountsFromConfig(c.RawAccounts)
	c.SecretKey = []byte(c.RawSecretKey)

	if c.UseAuth && len(c.SecretKey) == 0 {
		return fmt.Errorf("config: secret key is mandatory if using auth")
	}

	return nil
}

func (c *Config) IsEmailsEnabled() bool {
	return len(c.SmtpHost) > 0
}

func (c *Config) IsAuthAllowed(email string) bool {
	return c.Accounts.HasEmail(email)
}
