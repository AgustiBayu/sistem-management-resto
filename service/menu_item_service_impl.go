package service

import (
	"SistemManagementResto/exception"
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
	"SistemManagementResto/repository"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type MenuItemServiceImpl struct {
	MenuItemRepository repository.MenuItemRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewMenuItemService(menuItemRepository repository.MenuItemRepository, DB *sql.DB, validate *validator.Validate) MenuItemService {
	return &MenuItemServiceImpl{
		MenuItemRepository: menuItemRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *MenuItemServiceImpl) Create(ctx context.Context, request web.MenuItemCreateRequest) web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPenambahanItemMenu, err := time.Parse("02-01-2006", request.TanggalPenambahanItemMenu)
	helper.PanicIfError(err)
	menuItem := domain.MenuItem{
		Name:                      request.Name,
		Deskripsi:                 request.Name,
		Harga:                     request.Harga,
		TanggalPenambahanItemMenu: TanggalPenambahanItemMenu,
	}

	menuItem = service.MenuItemRepository.Save(ctx, tx, menuItem)
	return helper.ToMenuItemResponse(menuItem)
}
func (service *MenuItemServiceImpl) FindAll(ctx context.Context) []web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	menuItem := service.MenuItemRepository.FindAll(ctx, tx)
	return helper.ToMenuItemResponses(menuItem)
}
func (service *MenuItemServiceImpl) FindById(ctx context.Context, requestId int) web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToMenuItemResponse(menuItem)
}
func (service *MenuItemServiceImpl) Update(ctx context.Context, request web.MenuItemUpdateRequest) web.MenuItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	TanggalPenambahanItemMenu, err := time.Parse("02-01-2006", request.TanggalPenambahanItemMenu)
	helper.PanicIfError(err)
	err = helper.ValidateTanggalBaru(menuItem.TanggalPenambahanItemMenu, TanggalPenambahanItemMenu)
	helper.PanicIfError(err)

	menuItem.Name = request.Name
	menuItem.Deskripsi = request.Deskripsi
	menuItem.Harga = request.Harga
	menuItem.TanggalPenambahanItemMenu = TanggalPenambahanItemMenu
	menuItem = service.MenuItemRepository.Update(ctx, tx, menuItem)
	return helper.ToMenuItemResponse(menuItem)
}
func (service *MenuItemServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.MenuItemRepository.Delete(ctx, tx, menuItem)
}
