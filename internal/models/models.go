package models

type InputSignUp struct {
	ID       int    `json:"-"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type InputSignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
