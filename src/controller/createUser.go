package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("invalid json body")
	c.JSON(err.Code, err)
}
