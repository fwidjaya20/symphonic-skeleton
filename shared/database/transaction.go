package database

import (
	"github.com/fwidjaya20/symphonic-skeleton/shared/context"
	"github.com/labstack/echo/v4"
)

type TxCallback = func(e echo.Context) error

func RunInTransaction(c *context.SymphonicContext, callback TxCallback) error {
	tx := Gorm().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	c.Database = tx

	if err := callback(c); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
