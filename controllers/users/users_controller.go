package users

import (
	"github.com/Lozovoi-Rodion/bookstore_users-api/domain/users"
	"github.com/Lozovoi-Rodion/bookstore_users-api/services"
	"github.com/Lozovoi-Rodion/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			"invalid json body",
			http.StatusBadRequest,
			"bad_request",
		}
		c.JSON(restErr.Status, restErr)
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle user creation error
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
