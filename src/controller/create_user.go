package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leohenriquet/crud-golang/src/configuration/logger"
	"github.com/leohenriquet/crud-golang/src/configuration/validation"
	"github.com/leohenriquet/crud-golang/src/controller/model/request"
	"github.com/leohenriquet/crud-golang/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"))
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user request", err,
			zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Name:  UserRequest.Name,
		Email: UserRequest.Email,
		Age:   UserRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, response)
}
