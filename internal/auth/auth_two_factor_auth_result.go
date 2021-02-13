package auth

type authTwoFactorAuthResult struct {
	Repeat bool   `json:"repeat"`
	Error  string `json:"error"`
	Token  string `json:"token"`
}
