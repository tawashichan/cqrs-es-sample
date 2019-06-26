package service

import (
	"cqrs-es/domain/event/userEvent"
	"cqrs-es/domain/model/userModel"
)

type UserService struct {
	userEventRepository userModel.IUserEventRepository
}

func NewUserService(userRepository userModel.IUserEventRepository) UserService {
	return UserService{
		userEventRepository: userRepository,
	}
}

func (s UserService) UpdateUser(e *userEvent.UpdateEmail) error {
	return s.userEventRepository.UpdateUser(e)
}
