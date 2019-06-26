package onMemory

import (
	"cqrs-es/domain/event/userEvent"
	"cqrs-es/domain/model/userModel"
)

type userEventRepository struct{}

var updateUserArray = []*userEvent.UpdateEmail{}

func InitializeUserEventRepository() {
	updateUserArray = []*userEvent.UpdateEmail{}
}

func NewUserEventRepository() userModel.IUserEventRepository {
	return userEventRepository{}
}

func (userEventRepository) UpdateUser(e *userEvent.UpdateEmail) error {
	updateUserArray = append(updateUserArray, e)
	return nil
}
