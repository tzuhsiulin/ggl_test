package utils

import "database/sql"

func IsDataNotFoundErr(err error) bool {
	if err == sql.ErrNoRows {
		return true
	}
	return false
}
