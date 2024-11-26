package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err"
	"github.com/leohenriquet/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, errors: %s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(UserRequest)
}
