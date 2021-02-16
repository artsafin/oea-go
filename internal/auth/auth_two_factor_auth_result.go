package auth

type authTwoFactorAuthResult struct {
	Repeat    bool   `json:"repeat"`
	Error     string `json:"error"`
	Token     string `json:"token"`
	ReturnUrl string `json:"return"`
}

func new2FaError(err error, repeat bool) authTwoFactorAuthResult {
	return authTwoFactorAuthResult{
		Error:  err.Error(),
		Repeat: repeat,
	}
}

func new2FaSuccess(token, retUrl string) authTwoFactorAuthResult {
	return authTwoFactorAuthResult{
		Token:     token,
		ReturnUrl: retUrl,
	}
}
