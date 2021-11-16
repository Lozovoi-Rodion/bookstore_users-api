package mysql_utils

import (
	"github.com/Lozovoi-Rodion/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", nil)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request", nil)
}
