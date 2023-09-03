package tx

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Write(ctx context.Context, innerFunc func(tx *sql.Tx) error) error {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := innerFunc(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
