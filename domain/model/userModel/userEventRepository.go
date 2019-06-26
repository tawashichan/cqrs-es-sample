package userModel

import "cqrs-es/domain/event/userEvent"

type IUserEventRepository interface {
	UpdateUser(e *userEvent.UpdateEmail) error
}
