package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"gopkg.in/yaml.v3"
	"notes/internal/users"
	"os"
)

type Service struct {
	db        *sql.DB
	usersData *users.Data
}

func InitServices(database *sql.DB) *Service {
	return &Service{
		db:        database,
		usersData: users.NewData(database),
	}
}

type secretKey struct {
	Value string `yaml:"SECRET_KEY"`
}

func getSecretKey() (string, error) {
	yamlFile, err := os.ReadFile("./cmd/config.yaml")

	if err != nil {
		return "", err
	}

	var config *secretKey

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		return "", err
	}

	return config.Value, nil
}

func (s *Service) HashPassword(password string) (string, error) {
	sk, err := getSecretKey()

	if err != nil {
		return "", err
	}

	h := hmac.New(sha512.New, []byte(sk))
	_, err = h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash), nil
}
