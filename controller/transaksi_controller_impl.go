package controller

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/web"
	"SistemManagementResto/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
}

func NewTransaksiController(transaksiService service.TransaksiService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
	}
}

func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiCreateRequest := web.TransaksiCreateRequest{}
	helper.ReadRequestBody(request, &transaksiCreateRequest)

	transaksiResponse := controller.TransaksiService.Create(request.Context(), transaksiCreateRequest)
	liks := helper.CreateLinksForItem(transaksiResponse.Id, "transaksis")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  transaksiResponse,
		Links: liks,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *TransaksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiResponse := controller.TransaksiService.FindAll(request.Context())
	var ids []int
	for _, transaksi := range transaksiResponse {
		ids = append(ids, transaksi.Id)
	}
	liks := helper.CreateLinksForItems(ids, "transaksis")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  transaksiResponse,
		Links: liks,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *TransaksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiId := params.ByName("transaksiId")
	id, err := strconv.Atoi(transaksiId)
	helper.PanicIfError(err)

	transaksiResponse := controller.TransaksiService.FindById(request.Context(), id)
	liks := helper.CreateLinksForItem(transaksiResponse.Id, "transaksis")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  transaksiResponse,
		Links: liks,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *TransaksiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiUpdateRequest := web.TransaksiUpdateRequest{}
	helper.ReadRequestBody(request, &transaksiUpdateRequest)
	transaksiId := params.ByName("transaksiId")
	id, err := strconv.Atoi(transaksiId)
	helper.PanicIfError(err)

	transaksiUpdateRequest.Id = id
	transaksiResponse := controller.TransaksiService.Update(request.Context(), transaksiUpdateRequest)
	liks := helper.CreateLinksForItem(transaksiResponse.Id, "transaksis")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  transaksiResponse,
		Links: liks,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *TransaksiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiId := params.ByName("transaksiId")
	id, err := strconv.Atoi(transaksiId)
	helper.PanicIfError(err)

	controller.TransaksiService.Delete(request.Context(), id)
	webResponse := web.WebRespon{
		Code: 200,
	}
	helper.WriteResponBody(writer, webResponse)
}
