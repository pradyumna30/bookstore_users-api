package user_controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pradyumna30/bookstore_users-api/domain/users"
	"github.com/pradyumna30/bookstore_users-api/services"
	"github.com/pradyumna30/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid user id")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// If we reach here, entire business logic will be in service

	result, SaveErr := services.CreateUser(user)
	if SaveErr != nil {
		c.JSON(SaveErr.Status, SaveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
