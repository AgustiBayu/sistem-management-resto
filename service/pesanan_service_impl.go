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

type PesananServiceImpl struct {
	PesananRespository repository.PesananRepository
	UserRepository     repository.UserRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewPesananService(pesananRepository repository.PesananRepository, userRepository repository.UserRepository,
	DB *sql.DB, validate *validator.Validate) PesananService {
	return &PesananServiceImpl{
		PesananRespository: pesananRepository,
		UserRepository:     userRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *PesananServiceImpl) Create(ctx context.Context, request web.PesananCreateRequest) web.PesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	TanggalPembuatanPesanan, err := time.Parse("02-01-2006", request.TanggalPembuatanPesanan)
	helper.PanicIfError(err)
	pesanan := domain.Pesanan{
		UserId:                  request.UserId,
		TotalHarga:              request.TotalHarga,
		Status:                  domain.StatusPesanan(request.Status),
		TanggalPembuatanPesanan: TanggalPembuatanPesanan,
	}
	pesanan = service.PesananRespository.Save(ctx, tx, pesanan)
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) FindAll(ctx context.Context) []web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, user := service.PesananRespository.FindAll(ctx, tx)
	return helper.ToPesananResponses(pesanan, user)

}
func (service *PesananServiceImpl) FindById(ctx context.Context, requestId int) web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, user, err := service.PesananRespository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) Update(ctx context.Context, request web.PesananUpdateRequest) web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPembuatanPesanan, err := time.Parse("02-01-2006", request.TanggalPembuatanPesanan)
	helper.PanicIfError(err)
	pesanan, _, err := service.PesananRespository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = helper.ValidateTanggalBaru(pesanan.TanggalPembuatanPesanan, TanggalPembuatanPesanan)
	helper.PanicIfError(err)
	pesanan.UserId = request.Id
	pesanan.TotalHarga = request.TotalHarga
	pesanan.Status = domain.StatusPesanan(request.Status)
	pesanan.TanggalPembuatanPesanan = TanggalPembuatanPesanan

	pesanan = service.PesananRespository.Update(ctx, tx, pesanan)
	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, _, err := service.PesananRespository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PesananRespository.Delete(ctx, tx, pesanan)
}
