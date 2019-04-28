package onMemory

import (
	"cqrs-es/domain/model/userModel"
)

type userEventRepository struct{}

var updateUserArray = []*userModel.UpdateEmail{}

func InitializeUserEventRepository() {
	updateUserArray = []*userModel.UpdateEmail{}
}

func NewUserEventRepository() userModel.IUserEventRepository {
	return userEventRepository{}
}

func (userEventRepository) UpdateUser(e *userModel.UpdateEmail) error {
	updateUserArray = append(updateUserArray, e)
	return nil
}
