package userModel

import "cqrs-es/domain/model/common"

type IUserRepository interface {
	Get(id common.UserUUID)
}
