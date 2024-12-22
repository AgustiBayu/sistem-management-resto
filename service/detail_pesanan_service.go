package service

import (
	"SistemManagementResto/model/web"
	"context"
)

type DetailPesananService interface {
	Create(ctx context.Context, request web.DetailPesananCreateRequest) web.DetailPesananResponse
	FindAll(ctx context.Context) []web.DetailPesananResponse
	FindById(ctx context.Context, requestId int) web.DetailPesananResponse
	Update(ctx context.Context, request web.DetailPesananUpdateRequest) web.DetailPesananResponse
	Delete(ctx context.Context, requestId int)
}
