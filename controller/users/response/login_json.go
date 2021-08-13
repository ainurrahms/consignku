package response

type Login struct {
	Token string `json:"token"`
}

func FromDomainToLoginResponse(token string) Login {
	return Login{
		Token: token,
	}
}
