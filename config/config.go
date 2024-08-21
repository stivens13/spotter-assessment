package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *DBConfig) GetDSN() string {
    return fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
        c.Host,
        c.User,
        c.Password,
        c.Database,
        c.Port,
    )
}

func InitConfig() *Config {
	if err := godotenv.Load("postgres.env"); err != nil {
		log.Fatal("Error loading postgres.env file")
	}

	return &Config{
		DBConfig: &DBConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Database: os.Getenv("POSTGRES_DB"),
		},
	}
}
