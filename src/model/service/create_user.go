package service

import (
	"crud-go/src/configuration/logger"
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"fmt"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
