package service

import (
	"SistemManagementResto/model/web"
	"context"
)

type MenuItemService interface {
	Create(ctx context.Context, request web.MenuItemCreateRequest) web.MenuItemResponse
	FindAll(ctx context.Context) []web.MenuItemResponse
	FindById(ctx context.Context, requestId int) web.MenuItemResponse
	Update(ctx context.Context, request web.MenuItemUpdateRequest) web.MenuItemResponse
	Delete(ctx context.Context, requestId int)
}
