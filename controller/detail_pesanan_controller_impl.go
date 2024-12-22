package controller

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/web"
	"SistemManagementResto/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type DetailPesananControllerImpl struct {
	DetailPesananServie service.DetailPesananService
}

func NewDetailPesananController(detailPesananService service.DetailPesananService) DetailPesananController {
	return &DetailPesananControllerImpl{
		DetailPesananServie: detailPesananService,
	}
}

func (controller *DetailPesananControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananCreateRequest := web.DetailPesananCreateRequest{}
	helper.ReadRequestBody(request, &detailPesananCreateRequest)

	detailPesananResponse := controller.DetailPesananServie.Create(request.Context(), detailPesananCreateRequest)
	links := helper.CreateLinksForItem(detailPesananResponse.Id, "detaiPesanans")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  detailPesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananResponse := controller.DetailPesananServie.FindAll(request.Context())
	var ids []int
	for _, detailPesanan := range detailPesananResponse {
		ids = append(ids, detailPesanan.Id)
	}
	links := helper.CreateLinksForItems(ids, "detailPesanans")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  detailPesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)

	detailPesananResponse := controller.DetailPesananServie.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(detailPesananResponse.Id, "detailPesananId")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  detailPesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananUpdateRequest := web.DetailPesananUpdateRequest{}
	helper.ReadRequestBody(request, &detailPesananUpdateRequest)
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)
	detailPesananUpdateRequest.Id = id

	detailPesananResponse := controller.DetailPesananServie.Update(request.Context(), detailPesananUpdateRequest)
	links := helper.CreateLinksForItem(detailPesananResponse.Id, "detailPesananId")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  detailPesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)

	controller.DetailPesananServie.Delete(request.Context(), id)
	webResponse := web.WebRespon{
		Code: 200,
	}
	helper.WriteResponBody(writer, webResponse)
}
