package users

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type Data struct {
	db *sql.DB
}

func NewData(database *sql.DB) *Data {
	return &Data{db: database}
}

func (d *Data) CreateUser(email, hashedPassword string) (*User, error) {
	query := `
		INSERT INTO users (email, hashed_password)
		VALUES ($1, $2)
		RETURNING id, email, created_at
	`
	var newUser User
	err := d.db.QueryRow(query, email, hashedPassword).Scan(&newUser.Id, &newUser.Email, &newUser.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "unique constraint") {
			return nil, errors.New("user already exists")
		}
		return nil, errors.New("error creating user")
	}

	return &newUser, nil
}
