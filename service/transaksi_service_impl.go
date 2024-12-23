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

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	PesananRepository   repository.PesananRepository
	UserRepository      repository.UserRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, pesananRepository repository.PesananRepository,
	userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		PesananRepository:   pesananRepository,
		UserRepository:      userRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *TransaksiServiceImpl) Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, request.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if request.JumlahBayar < pesanan.TotalHarga {
		panic("jumlah bayar kurang dari total")
	}
	Kembalian := request.JumlahBayar - pesanan.TotalHarga

	TanggalTransaksi, err := time.Parse("02-01-2006", request.TanggalTransaksi)
	helper.PanicIfError(err)
	transaksi := domain.Transaksi{
		PesananId:        request.PesananId,
		MetodePembayaran: domain.StatusMetodePembayaran(request.MetodePembayaran),
		JumlahBayar:      request.JumlahBayar,
		Kembalian:        Kembalian,
		TanggalTransaksi: TanggalTransaksi,
	}
	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi)
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToTransaksiResponse(transaksi, pesanan, user)
}
func (service *TransaksiServiceImpl) FindAll(ctx context.Context) []web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	transaksi, pesanan, user := service.TransaksiRepository.FindAll(ctx, tx)
	return helper.ToTransaksiResponses(transaksi, pesanan, user)
}
func (service *TransaksiServiceImpl) FindById(ctx context.Context, requestId int) web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	transaksi, pesanan, user, err := service.TransaksiRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToTransaksiResponse(transaksi, pesanan, user)
}
func (service *TransaksiServiceImpl) Update(ctx context.Context, request web.TransaksiUpdateRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalTransaksi, err := time.Parse("02-01-2006", request.TanggalTransaksi)
	helper.PanicIfError(err)
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, request.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if request.JumlahBayar < pesanan.TotalHarga {
		panic("jumlah bayar kurang dari total")
	}
	Kembalian := request.JumlahBayar - pesanan.TotalHarga

	transaksi, _, _, err := service.TransaksiRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = helper.ValidateTanggalBaru(transaksi.TanggalTransaksi, TanggalTransaksi)
	helper.PanicIfError(err)
	transaksi.PesananId = request.PesananId
	transaksi.MetodePembayaran = domain.StatusMetodePembayaran(request.MetodePembayaran)
	transaksi.JumlahBayar = request.JumlahBayar
	transaksi.Kembalian = Kembalian
	transaksi.TanggalTransaksi = TanggalTransaksi

	transaksi = service.TransaksiRepository.Update(ctx, tx, transaksi)
	return helper.ToTransaksiResponse(transaksi, pesanan, user)
}
func (service *TransaksiServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	transaksi, _, _, err := service.TransaksiRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.TransaksiRepository.Delete(ctx, tx, transaksi)
}
