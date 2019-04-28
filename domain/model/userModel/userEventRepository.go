package userModel

type IUserEventRepository interface {
	UpdateUser(e *UpdateEmail) error
}
