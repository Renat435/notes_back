package users

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
