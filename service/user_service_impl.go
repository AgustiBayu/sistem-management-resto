package service

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
	"SistemManagementResto/repository"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository,
	DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalBuatAkun, err := time.Parse("02-01-2006", request.TanggalBuatAkun)
	helper.PanicIfError(err)
	hasPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	user := domain.User{
		Name:            request.Name,
		Email:           request.Email,
		Password:        string(hasPass),
		Pengguna:        domain.StatusPengguna(request.Pengguna),
		TanggalBuatAkun: TanggalBuatAkun,
	}
	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}
func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{}, errors.New("email is not valid")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("password is not valid")
	}
	return helper.ToUserResponse(user), nil
}
