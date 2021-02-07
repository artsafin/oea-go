package common

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func isValueEmpty(v interface{}) bool {
	switch v.(type) {
	case string:
		return v == ""
	case uint:
		return v == 0
	case int:
		return v == 0
	}

	return true
}

type ConfigTags struct {
	Name        string
	IsMandatory bool
}

type Config struct {
	BaseUri            string `oea:"base_uri,mandatory"`
	ApiTokenOf         string `oea:"api_token_of,mandatory"`
	ApiTokenEm         string `oea:"api_token_em,mandatory"`
	DocIdOf            string `oea:"doc_id_of,mandatory"`
	DocIdEm            string `oea:"doc_id_em,mandatory"`
	secretKeyStr       string `oea:"auth_key,mandatory"`
	SecretKey          []byte
	AuthAllowedDomains string `oea:"auth_allowed_domains,"`
	rawAccounts        string `oea:"auth_accounts,"`
	SmtpHost           string `oea:"smtp_host,mandatory"`
	SmtpUser           string `oea:"smtp_user,mandatory"`
	SmtpPass           string `oea:"smtp_pass,mandatory"`
	SmtpPort           uint   `oea:"smtp_port,mandatory"`
	SecurePort         uint   `oea:"secure_port,mandatory"`
	InsecurePort       uint   `oea:"insecure_port,mandatory"`
	TlsCert            string `oea:"tls_cert,"`
	TlsKey             string `oea:"tls_key,"`
	UseAuth            bool   `oea:"use_auth"`
	Accounts           Accounts
	AppVersion         string
	BotToken           string
}

func NewConfig(appVersion string) Config {
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
		"\033[33mUser Auth Restrictions:\033[0m %v; %v",
		"\033[33mSmtp:\033[0m %v:%v (%v)",
		"",
	}
	params := []interface{}{
		c.UseAuth,
		boolStr(c.TlsCert != "", "configured", "disabled"),
		c.DocIdOf,
		c.DocIdEm,
		c.AuthAllowedDomains,
		c.Accounts,
		c.SmtpHost,
		c.SmtpPort,
		boolStr(c.SmtpPass != "", "auth enabled", "auth disabled"),
	}
	_, _ = wr.Write([]byte(fmt.Sprintf(strings.Join(lines, "\n"), params...)))
}

func (c *Config) CastStringValueToConfigType(fieldName string, val string) interface{} {
	typ := reflect.TypeOf(c)
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic(fmt.Sprintf("config: CastStringValueToConfigType: field %s was not found", fieldName))
	}

	switch field.Type.Kind() {
	case reflect.Uint:
		newValUint, convErr := strconv.ParseUint(val, 10, 64)
		if convErr != nil {
			return nil
		}
		return uint(newValUint)
	case reflect.Bool:
		boolVal := val != "0" && val != "false" && val != "no"
		return boolVal
	default:
		return val
	}
}

func parseTags(fieldType reflect.StructField, tag string) ConfigTags {
	tags := ConfigTags{
		IsMandatory: false,
		Name:        fieldType.Name,
	}

	tagVals := strings.Split(tag, ",")

	if len(tagVals) == 1 {
		tags.Name = tagVals[0]
	} else if len(tagVals) > 1 {
		tags.Name = tagVals[0]
		tags.IsMandatory = tagVals[1] == "mandatory"
	}

	return tags
}

func (c *Config) ForeachKey(mapFn func(name string, val interface{}, configTags ConfigTags) (interface{}, bool)) {
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	for i := 0; i < val.Elem().NumField(); i++ {
		fieldVal := val.Elem().Field(i)
		fieldType := typ.Elem().Field(i)
		if !fieldVal.CanInterface() {
			continue
		}
		tags := parseTags(fieldType, fieldType.Tag.Get("oea"))

		newVal, shouldContinue := mapFn(fieldType.Name, fieldVal.Interface(), tags)

		if newVal != nil {
			if !fieldVal.CanSet() {
				log.Printf("config: %v: field is not settable. Skipping\n", fieldType.Name)
				continue
			}
			reflNewVal := reflect.ValueOf(newVal)
			if fieldVal.Type() != reflNewVal.Type() {
				log.Printf("config: %v: type mismatch: %v is not assignable to %v. Skipping\n", fieldType.Name, reflNewVal.Kind(), fieldVal.Kind())
				continue
			}

			fieldVal.Set(reflNewVal)
		}

		if !shouldContinue {
			break
		}
	}
}

func (c *Config) IsTLS() bool {
	return c.TlsCert != ""
}

func (c *Config) getEmptyMandatoryFields() []string {
	emptyFields := make([]string, 0)

	c.ForeachKey(func(fieldName string, val interface{}, tags ConfigTags) (interface{}, bool) {
		if !tags.IsMandatory {
			return nil, true
		}

		if isValueEmpty(val) {
			emptyFields = append(emptyFields, tags.Name)
		}
		return nil, true
	})

	return emptyFields
}

func (c *Config) PrepareAndValidate() error {
	if empty := c.getEmptyMandatoryFields(); len(empty) > 0 {
		return fmt.Errorf("config validation: not all config parameters are set: %v", empty)
	}

	c.Accounts = newAccountsFromConfig(c.rawAccounts)
	c.SecretKey = []byte(c.secretKeyStr)

	// TODO: fix
	c.BotToken = "1605826773:AAFdPBF-gAXsK97qoBOy_d-QSgkJsYsodvw"

	return nil
}

func extractEmailDomain(email string) string {
	if emailDomainIdx := strings.Index(email, "@"); emailDomainIdx >= 0 {
		return email[emailDomainIdx+1:]
	}
	return email
}

func (c *Config) isAuthAllowedForDomain(email string) bool {
	domains := strings.Split(c.AuthAllowedDomains, ",")

	emailDomain := extractEmailDomain(email)

	for _, domain := range domains {
		if strings.TrimSpace(domain) == emailDomain {
			return true
		}
	}
	return false
}

func (c *Config) isAuthAllowedForEmail(email string) bool {
	return c.Accounts.HasEmail(email)
}

func (c *Config) IsAuthAllowed(email string) bool {
	return c.isAuthAllowedForEmail(email) || c.isAuthAllowedForDomain(email)
}
