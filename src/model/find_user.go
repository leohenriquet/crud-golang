package model

import "github.com/leohenriquet/crud-golang/src/configuration/rest_err"

func (*UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	return nil, nil
}