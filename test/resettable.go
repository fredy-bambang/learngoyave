package test

import (
	"context"

	"github.com/fredy-bambang/learngoyave/database/sqlboiler/model"
	"goyave.dev/goyave/v4/database"
)

func resetTable() error {
	db, _ := database.Conn().DB()
	_, err := model.Users().DeleteAll(context.Background(), db)

	if err != nil {
		return err
	}

	return nil
}
