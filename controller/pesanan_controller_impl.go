package controller

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/web"
	"SistemManagementResto/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PesananControllerImpl struct {
	PesananService service.PesananService
}

func NewPesananController(pesananService service.PesananService) PesananController {
	return &PesananControllerImpl{
		PesananService: pesananService,
	}
}

func (controller *PesananControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pesananCreateRequest := web.PesananCreateRequest{}
	helper.ReadRequestBody(request, &pesananCreateRequest)

	pesananResponse := controller.PesananService.Create(request.Context(), pesananCreateRequest)
	links := helper.CreateLinksForItem(pesananResponse.Id, "pesanan")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  pesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *PesananControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pesananResponse := controller.PesananService.FindAll(request.Context())
	var ids []int
	for _, pesanan := range pesananResponse {
		ids = append(ids, pesanan.Id)
	}
	links := helper.CreateLinksForItems(ids, "pesanans")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  pesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *PesananControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pesananId := params.ByName("pesananId")
	id, err := strconv.Atoi(pesananId)
	helper.PanicIfError(err)
	pesananResponse := controller.PesananService.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(pesananResponse.Id, "pesanans")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  pesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *PesananControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pesananUpdateRequest := web.PesananUpdateRequest{}
	helper.ReadRequestBody(request, &pesananUpdateRequest)
	pesananId := params.ByName("pesananId")
	id, err := strconv.Atoi(pesananId)
	helper.PanicIfError(err)

	pesananUpdateRequest.Id = id
	pesananResponse := controller.PesananService.Update(request.Context(), pesananUpdateRequest)
	links := helper.CreateLinksForItem(pesananResponse.Id, "pesanans")
	webResponse := web.WebResponses{
		Code:  200,
		Data:  pesananResponse,
		Links: links,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *PesananControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pesananId := params.ByName("pesananId")
	id, err := strconv.Atoi(pesananId)
	helper.PanicIfError(err)

	controller.PesananService.Delete(request.Context(), id)
	webResponse := web.WebRespon{
		Code: 200,
	}
	helper.WriteResponBody(writer, webResponse)
}
