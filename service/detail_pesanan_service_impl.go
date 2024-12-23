package service

import (
	"SistemManagementResto/exception"
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
	"SistemManagementResto/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type DetailPesananServiceImpl struct {
	DetailPesananRepository repository.DetailPesananRespository
	PesananRepository       repository.PesananRepository
	UserRepository          repository.UserRepository
	MenuItemRepository      repository.MenuItemRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewDetailPesananService(detailPesananRepository repository.DetailPesananRespository,
	pesananRepository repository.PesananRepository, userRepository repository.UserRepository,
	menuItemRepository repository.MenuItemRepository,
	DB *sql.DB, validate *validator.Validate) DetailPesananService {
	return &DetailPesananServiceImpl{
		DetailPesananRepository: detailPesananRepository,
		PesananRepository:       pesananRepository,
		UserRepository:          userRepository,
		MenuItemRepository:      menuItemRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *DetailPesananServiceImpl) Create(ctx context.Context, request web.DetailPesananCreateRequest) web.DetailPesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	detailPesanan := domain.DetailPesanan{
		PesananId:         request.PesananId,
		MenuItemId:        request.MenuItemId,
		JumlahItemPesanan: request.JumlahItemPesanan,
		HargaItem:         request.HargaItem,
	}
	detailPesanan = service.DetailPesananRepository.Save(ctx, tx, detailPesanan)
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, detailPesanan.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, detailPesanan.MenuItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToDetailPesananResponse(detailPesanan, pesanan, user, menuItem)
}
func (service *DetailPesananServiceImpl) FindAll(ctx context.Context) []web.DetailPesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detailPesanan, pesanan, user, menuItem := service.DetailPesananRepository.FindAll(ctx, tx)
	return helper.ToDetailPesananResponses(detailPesanan, pesanan, user, menuItem)
}
func (service *DetailPesananServiceImpl) FindById(ctx context.Context, requestId int) web.DetailPesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	detailPesanan, pesanan, user, menuItem, err := service.DetailPesananRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToDetailPesananResponse(detailPesanan, pesanan, user, menuItem)
}
func (service *DetailPesananServiceImpl) Update(ctx context.Context, request web.DetailPesananUpdateRequest) web.DetailPesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	detailPesanan, _, _, _, err := service.DetailPesananRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, detailPesanan.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	menuItem, err := service.MenuItemRepository.FindById(ctx, tx, detailPesanan.MenuItemId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	detailPesanan.PesananId = request.PesananId
	detailPesanan.MenuItemId = request.MenuItemId
	detailPesanan.JumlahItemPesanan = request.JumlahItemPesanan
	detailPesanan.HargaItem = request.HargaItem
	detailPesanan = service.DetailPesananRepository.Update(ctx, tx, detailPesanan)
	return helper.ToDetailPesananResponse(detailPesanan, pesanan, user, menuItem)
}
func (service *DetailPesananServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detailPesanan, _, _, _, err := service.DetailPesananRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.DetailPesananRepository.Delete(ctx, tx, detailPesanan)
}
