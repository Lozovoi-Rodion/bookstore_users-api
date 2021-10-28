package users

import (
	"fmt"
	"github.com/Lozovoi-Rodion/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Lozovoi-Rodion/bookstore_users-api/utils/date"
	"github.com/Lozovoi-Rodion/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "users.users_UN"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}
