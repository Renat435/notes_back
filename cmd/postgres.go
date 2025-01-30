package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Name     string `yaml:"POSTGRES_DB"`
	User     string `yaml:"POSTGRES_USER"`
	Password string `yaml:"POSTGRES_PASSWORD"`
	Port     string `yaml:"PORT"`
	Host     string `yaml:"HOST"`
}

func PostgresConnection() (*sql.DB, error) {
	config := getDatabaseConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDatabaseConfig() Config {
	yamlFile, err := os.ReadFile("./cmd/config.yaml")

	if err != nil {
		log.Fatalf("Read config file error: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatalf("Unmarshal config file error: %v", err)
	}

	return config
}
