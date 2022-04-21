package schema

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Login
}
