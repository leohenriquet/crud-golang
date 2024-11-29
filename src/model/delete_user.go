package model

import "github.com/leohenriquet/crud-golang/src/configuration/rest_err"

func (*UserDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}
