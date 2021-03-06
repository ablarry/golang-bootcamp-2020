package service

import (
	"database/sql"

	"github.com/ablarry/golang-bootcamp-2020/pkg/utils/errors"
	"github.com/lib/pq"
)

// Handler manage errors sql
func Handler(err error) error {
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrNotFound
		}
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			err = errors.ErrDuplicate
		}
	}
	return err
}	
