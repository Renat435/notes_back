package service

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
