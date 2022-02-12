package authtoken

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt"
	"github.com/pkg/errors"
	"oea-go/internal/auth/enc"
	"oea-go/internal/config"
	"strconv"
	"strings"
	"time"
)

const (
	JwtIssuerName = "oea-%s" // %s is app version
	JwtSubject    = "oea_users"
	JwtExpiration = time.Hour * 1
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
	Source    string
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
		Source:    source,
		Claims:    *claims,
		validator: newValidator(appVersion),
	}, nil
}

func GenerateTokenFirstFactor(appVersion string, audience string, authKey []byte, returnUrl string) (*jwt.Token, error) {
	return buildToken(
		appVersion,
		audience,
		authKey,
		time.Now().Add(JwtExpiration).Unix(),
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

func (tok *Token) Validate2FA(encSecretKey []byte, acc config.Account) (err error) {
	var key [32]byte
	copy(key[:], encSecretKey)

	//plain := fmt.Sprintf("%v:%v:%v:%v", r.userID, r.chatID, r.username, time.Now().Unix())

	fpHex, err := hex.DecodeString(tok.Claims.TwoFactorAuthFingerprint)
	if err != nil {
		return err
	}

	plain, err := enc.Decrypt(fpHex, key)
	if err != nil {
		return err
	}

	parts := strings.Split(string(plain), ":")
	if len(parts) != 4 {
		return errors.New("incorrect tff field: len is not 4")
	}

	tffTsInt, _ := strconv.ParseInt(parts[3], 10, 64)
	tffTs := time.Unix(tffTsInt, 0)
	now := time.Now()

	if tffTs.Before(now.Add(-JwtExpiration)) || tffTs.After(now.Add(JwtExpiration)) {
		return errors.Errorf("incorrect tff ts: %v %v", tffTs.Unix(), now.Unix())
	}

	if parts[2] != string(acc.ExternalUsername) {
		return errors.Errorf("incorrect tff username: %v %v", parts[2], acc.ExternalUsername)
	}

	return nil
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
