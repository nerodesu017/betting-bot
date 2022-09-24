package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func getDBUrl(host, port, user, pass, dbName string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbName)
}

var Pool *pgxpool.Pool
var err error

func init() {
	Pool, err = pgxpool.New(context.Background(),
		getDBUrl(os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME")))

	if err != nil {
		logrus.Fatalf("Couldn't connect to the DB; Error: %v", err)
	}

	logrus.Info("Connected to the database")
}
