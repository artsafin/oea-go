package common

import (
	"fmt"
	"log"
	"reflect"
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

type Config struct {
	BaseUri            string `oea:"base_uri,mandatory"`
	ApiTokenOf         string `oea:"api_token_of,mandatory"`
	ApiTokenEm         string `oea:"api_token_em,mandatory"`
	DocIdOf            string `oea:"doc_id_of,mandatory"`
	DocIdEm            string `oea:"doc_id_em,mandatory"`
	AuthKey            string `oea:"auth_key,mandatory"`
	AuthAllowedDomains string `oea:"auth_allowed_domains,"`
	AuthAllowedEmails  string `oea:"auth_allowed_emails,"`
	SmtpHost           string `oea:"smtp_host,mandatory"`
	SmtpUser           string `oea:"smtp_user,mandatory"`
	SmtpPass           string `oea:"smtp_pass,mandatory"`
	SmtpPort           uint   `oea:"smtp_port,mandatory"`
	SecurePort         uint   `oea:"secure_port,mandatory"`
	InsecurePort       uint   `oea:"insecure_port,mandatory"`
	TlsCert            string `oea:"tls_cert,"`
	TlsKey             string `oea:"tls_key,"`
	UseAuth            bool   `oea:"use_auth"`
}

func NewConfig() Config {
	return Config{
		BaseUri:      "https://coda.io/apis/v1beta1",
		SecurePort:   8443,
		InsecurePort: 8080,
		SmtpPort:     25,
		UseAuth:      true,
	}
}

func (c *Config) ForeachKey(mapFn func(name string, val interface{}, isMandatory bool, dstKind reflect.Kind) (interface{}, bool)) {
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	for i := 0; i < val.Elem().NumField(); i++ {
		fieldVal := val.Elem().Field(i)
		fieldType := typ.Elem().Field(i)
		if !fieldVal.CanInterface() {
			continue
		}
		tagVals := strings.SplitN(fieldType.Tag.Get("oea"), ",", 2)
		isMandatory := false
		fieldName := fieldType.Name
		if len(tagVals) == 1 {
			fieldName = tagVals[0]
		} else if len(tagVals) > 1 {
			fieldName = tagVals[0]
			isMandatory = tagVals[1] == "mandatory"
		}

		newVal, shouldContinue := mapFn(fieldName, fieldVal.Interface(), isMandatory, fieldVal.Kind())

		if newVal != nil {
			if !fieldVal.CanSet() {
				log.Printf("config: %v: field is not settable. Skipping\n", fieldName)
				continue
			}
			reflNewVal := reflect.ValueOf(newVal)
			if fieldVal.Type() != reflNewVal.Type() {
				log.Printf("config: %v: type mismatch: %v is not assignable to %v. Skipping\n", fieldName, reflNewVal.Kind(), fieldVal.Kind())
				continue
			}

			fieldVal.Set(reflNewVal)
		}

		if !shouldContinue {
			break
		}
	}
}

func (c Config) IsTLS() bool {
	return c.TlsCert != ""
}

func (c Config) getEmptyFields() []string {
	emptyFields := make([]string, 0)

	c.ForeachKey(func(fieldName string, val interface{}, isMandatory bool, dstKind reflect.Kind) (interface{}, bool) {
		if !isMandatory {
			return nil, true
		}

		if isValueEmpty(val) {
			emptyFields = append(emptyFields, fieldName)
		}
		return nil, true
	})

	return emptyFields
}

func (c Config) MustValidate() {
	if empty := c.getEmptyFields(); len(empty) > 0 {
		panic(fmt.Errorf("config validation: not all config parameters are set: %v", empty))
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
