package authtoken

import (
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt"
	"github.com/pkg/errors"
	"time"
)

const (
	JwtIssuerName = "oea-%s" // %s is app version
	JwtSubject    = "oea_users"
)

func issuerName(appVersion string) string {
	return fmt.Sprintf(JwtIssuerName, appVersion)
}

func newValidator(appVersion string) *jwt.Validator {
	return jwt.NewValidator(
		jwt.IssuerChecker(issuerName(appVersion)),
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

type JwtClaims struct {
	jwt.StandardClaims

	ReturnUrl                string `json:"ret,omitempty"`
	Version                  string `json:"ver,omitempty"`
	TwoFactorAuthFingerprint string `json:"tff,omitempty"`
}

func (c JwtClaims) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

type Token struct {
	Claims    JwtClaims
	validator *jwt.Validator
}

func (tok *Token) String() string {
	return fmt.Sprintf("[%+v]", tok.Claims)
}

func FromSource(appVersion string, authKey []byte, source string) (*Token, error) {
	if source == "" {
		return nil, errors.New("token is empty")
	}

	jwtTok, jwtErr := jwt.ParseAndVerifyString(source, newJwtSigner(authKey))

	if jwtErr != nil {
		return nil, fmt.Errorf("invalid jwt: %s", jwtErr)
	}

	claims := &JwtClaims{}
	jsonErr := json.Unmarshal(jwtTok.RawClaims(), claims)

	if jsonErr != nil {
		return nil, fmt.Errorf("invalid claims json: %s %q", jsonErr, jwtTok.RawClaims())
	}

	return &Token{
		Claims:    *claims,
		validator: newValidator(appVersion),
	}, nil
}

func GenerateTokenFirstFactor(appVersion string, audience string, authKey []byte, returnUrl string) (*jwt.Token, error) {
	return buildToken(
		appVersion,
		audience,
		authKey,
		time.Now().Add(time.Hour*1).Unix(),
		returnUrl,
		"",
	)
}

func GenerateTokenSecondFactor(firstFactorToken *Token, secondFactorFingerprint string, authKey []byte) (*jwt.Token, error) {
	audience, audErr := firstFactorToken.Email()
	if audErr != nil {
		return nil, audErr
	}

	return buildToken(
		firstFactorToken.Claims.Version,
		audience,
		authKey,
		firstFactorToken.ExpiresAt().Unix(),
		firstFactorToken.Claims.ReturnUrl,
		secondFactorFingerprint,
	)
}

func buildToken(appVersion string, audience string, authKey []byte, expiresAtTs int64, returnUrl string, fingerprint string) (*jwt.Token, error) {
	signer := newJwtSigner(authKey)

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuerName(appVersion),
			Subject:   JwtSubject,
			ExpiresAt: jwt.Timestamp(expiresAtTs),
			Audience:  []string{audience},
		},
		ReturnUrl:                returnUrl,
		Version:                  appVersion,
		TwoFactorAuthFingerprint: fingerprint,
	}

	return jwt.Build(signer, claims)
}

func (tok *Token) ValidateClaims() error {
	return tok.validator.Validate(&tok.Claims.StandardClaims)
}

func (tok *Token) Email() (string, error) {
	if len(tok.Claims.Audience) == 0 {
		return "", errors.Errorf("audience empty: %v", tok)
	}

	return tok.Claims.Audience[0], nil
}

func (tok *Token) ExpiresAt() time.Time {
	return tok.Claims.ExpiresAt.Time()
}

func newJwtSigner(authKey []byte) jwt.Signer {
	signer, _ := jwt.NewHS512(authKey)

	return signer
}

func CreateFromSourceAndValidate(appVersion string, authKey []byte, source string) (token *Token, err error) {
	tok, tokErr := FromSource(appVersion, authKey, source)
	if tokErr != nil {
		return nil, errors.Wrapf(tokErr, "token is invalid")
	}

	validationErr := tok.ValidateClaims()
	if validationErr != nil {
		return nil, errors.Wrapf(validationErr, "token claims invalid")
	}

	return tok, nil
}
