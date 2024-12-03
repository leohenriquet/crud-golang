package service

import (
	"github.com/leohenriquet/crud-golang/src/configuration/rest_err"
	"github.com/leohenriquet/crud-golang/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
