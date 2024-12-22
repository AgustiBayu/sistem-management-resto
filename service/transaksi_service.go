package service

import (
	"SistemManagementResto/model/web"
	"context"
)

type TransaksiService interface {
	Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse
	FindAll(ctx context.Context) []web.TransaksiResponse
	FindById(ctx context.Context, requestId int) web.TransaksiResponse
	Update(ctx context.Context, request web.TransaksiUpdateRequest) web.TransaksiResponse
	Delete(ctx context.Context, requestId int)
}
