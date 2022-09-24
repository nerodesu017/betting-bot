package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Fatalf("Error initializing .env file; error: %v", err)
	}
}
