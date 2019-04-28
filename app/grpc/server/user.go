package server

import (
	"context"
	"cqrs-es/app/grpc/protobuf"
	"cqrs-es/domain/model/userModel"
	"cqrs-es/domain/service"
)

type UserServer struct {
	userService service.UserService
}

func NewUserServer(userService service.UserService) UserServer {
	return UserServer{
		userService: userService,
	}
}

func (server *UserServer) UpdateEmail(ctx context.Context, req *user.UpdateUserEvent) (*user.ReturnOK, error) {
	event := &userModel.UpdateEmail{
		Email: req.Email,
	}
	if err := server.userService.UpdateUser(event); err != nil {
		return nil, err
	}
	return &user.ReturnOK{}, nil
}
