package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cristalhq/jwt"
	"oea-go/common"
	"time"
)

const (
	JwtIssuerName = "oea-%s" // %s is app version
	JwtSubject    = "oea_users"
)

func issuerName() string {
	return fmt.Sprintf(JwtIssuerName, AppVersion)
}

func newValidator() *jwt.Validator {
	return jwt.NewValidator(
		jwt.IssuerChecker(issuerName()),
		jwt.SubjectChecker(JwtSubject),
		func(claims *jwt.StandardClaims) error {
			if claims.ExpiresAt == 0 {
				return errors.New("jwt: exp claim is mandatory")
			}
			return nil
		},
		jwt.ExpirationTimeChecker(time.Now()),
	)
}

type Token struct {
	Claims    *jwt.StandardClaims
	validator *jwt.Validator
}

func (tok *Token) String() string {
	return fmt.Sprintf("[%+v]", *tok.Claims)
}

func VerifiedTokenFromSource(cfg common.Config, source string) (*Token, error) {
	jwtTok, jwtErr := jwt.ParseAndVerifyString(source, newJwtSigner(cfg))

	if jwtErr != nil {
		return nil, fmt.Errorf("invalid jwt: %s", jwtErr)
	}

	claims := &jwt.StandardClaims{}
	jsonErr := json.Unmarshal(jwtTok.RawClaims(), claims)

	if jsonErr != nil {
		return nil, fmt.Errorf("invalid claims json: %s %q", jsonErr, jwtTok.RawClaims())
	}

	return &Token{
		Claims:    claims,
		validator: newValidator(),
	}, nil
}

func GenerateToken(audience string, cfg common.Config) (string, error) {
	signer := newJwtSigner(cfg)

	claims := jwt.StandardClaims{
		Issuer: issuerName(),
		Subject: JwtSubject,
		ExpiresAt: jwt.Timestamp(time.Now().Add(time.Hour * 1).Unix()),
		Audience: []string{audience},
	}

	tok, err := jwt.Build(signer, claims)
	if err != nil {
		return "", err
	}

	return tok.InsecureString(), nil
}

func (tok *Token) ValidateClaims() error {
	return tok.validator.Validate(tok.Claims)
}

func newJwtSigner(config common.Config) jwt.Signer {
	signer, _ := jwt.NewHS512([]byte(config.AuthKey))

	return signer
}
