package userModel

import "cqrs-es/domain/model/common"

type UserReadModel struct {
	Id    common.UserUUID
	Email common.Email
	Name  UserName
}
