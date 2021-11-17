package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
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
	RawSecretKey string `envconfig:"secret_key"`
	RawAccounts  string `envconfig:"auth_accounts"`
	SmtpHost     string `envconfig:"smtp_host" required:"true"`
	SmtpUser     string `envconfig:"smtp_user" required:"true"`
	SmtpPass     string `envconfig:"smtp_pass" required:"true"`
	SmtpPort     uint   `envconfig:"smtp_port"`
	SecurePort   uint   `envconfig:"secure_port"`
	InsecurePort uint   `envconfig:"insecure_port"`
	TlsCert      string `envconfig:"tls_cert"`
	TlsKey       string `envconfig:"tls_key"`
	UseAuth      bool   `envconfig:"use_auth"`
	BotToken     string `envconfig:"bot_token"`
	FilesDir     string `envconfig:"files" required:"true"`
	IsDebug      bool   `envconfig:"debug"`
	StorageAddr  string
}

func NewDefaultConfig(appVersion string, storageAddr string) Config {
	return Config{
		SecurePort:   8443,
		InsecurePort: 8080,
		SmtpPort:     25,
		UseAuth:      true,
		AppVersion:   appVersion,
		StorageAddr:  storageAddr,
	}
}

func (c *Config) DumpNonSecretParameters(wr io.Writer) {
	yellowKV := func(header, value string) string {
		return fmt.Sprintf("\033[33m%v:\033[0m %v", header, value)
	}
	redKV := func(header, value string) string {
		return fmt.Sprintf("\033[31m%v:\033[0m %v", header, value)
	}

	lines := []string{
		redKV("Users", boolStr(c.UseAuth, c.Accounts.String(), "⚠ auth disabled")),
		redKV("TLS", boolStr(c.TlsCert != "", "configured", "⚠ disabled")),
		redKV("Debug mode", boolStr(c.IsDebug, "⚠ enabled", "disabled")),
		yellowKV("Storage", c.StorageAddr),
		yellowKV("Files", c.FilesDir),
		yellowKV("Document IDs", fmt.Sprintf("of=%v; em=%v", c.DocIdOf, c.DocIdEm)),
		yellowKV("SMTP", fmt.Sprintf(
			"%v (%v)",
			boolStr(c.SmtpHost != "", fmt.Sprintf("%v:%v", c.SmtpHost, c.SmtpPort), "email sending disabled"),
			boolStr(c.SmtpPass != "", "auth enabled", "auth disabled"),
		)),
		"",
	}

	_, _ = wr.Write([]byte(strings.Join(lines, "\n")))
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

	if c.UseAuth {
		if len(c.SecretKey) == 0 {
			return fmt.Errorf("config: secret key is mandatory if using auth")
		}

		if len(c.Accounts) == 0 {
			return fmt.Errorf("config: accounts must be preconfigured if using auth")
		}

		if len(c.BotToken) == 0 {
			return fmt.Errorf("config: bot token must be configured if using auth")
		}
	}

	if len(c.StorageAddr) == 0 {
		return errors.New("storage address is required")
	}

	return nil
}

func (c *Config) IsEmailsEnabled() bool {
	return len(c.SmtpHost) > 0
}

func (c *Config) IsAuthAllowed(email string) bool {
	return c.Accounts.HasEmail(email)
}
