package controller

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/web"
	"SistemManagementResto/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MenuItemControllerImpl struct {
	MenuItemService service.MenuItemService
}

func NewMenuItemController(menuItemService service.MenuItemService) MenuItemController {
	return &MenuItemControllerImpl{
		MenuItemService: menuItemService,
	}
}

func (controller *MenuItemControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuItemCreateRequest := web.MenuItemCreateRequest{}
	helper.ReadRequestBody(request, &menuItemCreateRequest)

	menuItemResponse := controller.MenuItemService.Create(request.Context(), menuItemCreateRequest)
	links := helper.CreateLinksForMenuItem(menuItemResponse)
	webResponse := web.WebResponses{
		Code:  200,
		Data:  menuItemResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *MenuItemControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuItemResponse := controller.MenuItemService.FindAll(request.Context())
	links := helper.CreateLinksForMenuItems(menuItemResponse)
	webResponse := web.WebResponses{
		Code:  200,
		Data:  menuItemResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *MenuItemControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuItemId := params.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	menuItemResponse := controller.MenuItemService.FindById(request.Context(), id)
	links := helper.CreateLinksForMenuItem(menuItemResponse)
	webResponse := web.WebResponses{
		Code:  200,
		Data:  menuItemResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *MenuItemControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuItemUpdateRequest := web.MenuItemUpdateRequest{}
	helper.ReadRequestBody(request, &menuItemUpdateRequest)
	menuItemId := params.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	menuItemUpdateRequest.Id = id
	menuItemResponse := controller.MenuItemService.Update(request.Context(), menuItemUpdateRequest)
	links := helper.CreateLinksForMenuItem(menuItemResponse)
	webResponse := web.WebResponses{
		Code:  200,
		Data:  menuItemResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *MenuItemControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	menuItemId := params.ByName("menuItemId")
	id, err := strconv.Atoi(menuItemId)
	helper.PanicIfError(err)

	controller.MenuItemService.Delete(request.Context(), id)
	webResponse := web.WebRespon{
		Code: 200,
	}
	helper.WriteResponBody(writer, webResponse)
}
