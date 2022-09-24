package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type UserStruct struct {
	Id             string
	Balance        int
	LastAllocation time.Time
}

func createUserTableIfNotExists() {
	Pool.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS public.users
(
    id character(18) NOT NULL,
    balance integer NOT NULL CHECK (balance > -1),
    "lastAllocation" date,
    PRIMARY KEY (id)
);
	`)
}

func InsertUser(userID string) error {
	if len(userID) != 18 {
		return errors.New(errors_userIDLength)
	}

	date := time.Now().AddDate(0, 0, -1) // set the date to yesterday, so use can get credits

	_, err := Pool.Exec(context.Background(), `
	INSERT INTO public.users(
	id, balance, "lastAllocation")
	VALUES ($1, $2, $3)
	ON CONFLICT (id) DO NOTHING;
	`, userID, 0, date.Format("2006-01-02"))

	if err != nil {
		logrus.Errorf("%v", err)
		return err
	}

	return nil
}

func SelectUser(userID string) (*UserStruct, error) {
	if len(userID) != 18 {
		return nil, errors.New(errors_userIDLength)
	}

	user := UserStruct{
		Id: userID,
	}
	err := Pool.QueryRow(context.Background(), `
	SELECT 	balance,
					"lastAllocation"
		FROM public.users
		WHERE id=$1
	`, userID).Scan(&user.Balance, &user.LastAllocation)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New(errors_noUserFound)
		}
		logrus.Errorf("%v", err)
		return nil, err
	}

	return &user, nil
}

func CreditUser(userID string, credits int) error {
	if len(userID) != 18 {
		return errors.New(errors_userIDLength)
	}

	now := time.Now().Format("2006-01-2")
	result, err := Pool.Exec(context.Background(), `
	UPDATE public.users
	SET balance=balance+$1,
			"lastAllocation"=$2
	WHERE id=$3
	AND "lastAllocation"<>$4
	`, credits, now, userID, now)

	if err != nil {
		logrus.Errorf("%v", err)
		return errors.New("internal error, contact bot owner")
	}

	if result.RowsAffected() == 0 {
		return errors.New("you've got your credits for the day")
	}

	return nil
}

func HasEnoughCredits(userID string, amount int64) bool {
	if len(userID) != 18 {
		return false
	}

	var hasEnough bool
	err := Pool.QueryRow(context.Background(), `
	SELECT EXISTS(
		SELECT COUNT(*)
		FROM public.users
		WHERE id=$1
		AND balance>=$2
	)
		`, userID, amount).Scan(&hasEnough)

	if err != nil || !hasEnough {
		logrus.Errorf("%v", err)
		return false
	}

	return true
}

func ChangeAndRetrieveBalance(userID string, amount int64) (int64, error) {
	if len(userID) != 18 {
		return 0, errors.New(errors_userIDLength)
	}

	var balance int64
	err := Pool.QueryRow(context.Background(), `
	UPDATE public.users
	SET balance=balance+$1
	WHERE id=$2
	RETURNING balance
	`, amount, userID).Scan(&balance)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func init() {
	createUserTableIfNotExists()
}
