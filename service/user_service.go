package service

import (
	"SistemManagementResto/model/web"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
}
