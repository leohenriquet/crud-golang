package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err/validation"
	"github.com/leohenriquet/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(UserRequest)
}
