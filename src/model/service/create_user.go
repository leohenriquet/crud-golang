package service

import (
	"fmt"

	"github.com/leohenriquet/crud-golang/src/configuration/logger"
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err"
	"github.com/leohenriquet/crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDoamin model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDoamin.EncryptPassword()

	fmt.Println(userDoamin.GetPassword())
	return nil
}
