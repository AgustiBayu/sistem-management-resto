package exception

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notfoundError(writer, request, err) {
		return
	}
	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Conten-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webRespone := web.WebResponse{
			Code: http.StatusBadRequest,
			Data: exception.Error(),
		}
		helper.WriteResponBody(writer, webRespone)
		return true
	} else {
		return false
	}
}

func notfoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Conten-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webRespone := web.WebResponse{
			Code: http.StatusNotFound,
			Data: exception.Error,
		}
		helper.WriteResponBody(writer, webRespone)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Conten-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webRespone := web.WebResponse{
		Code: http.StatusInternalServerError,
		Data: err,
	}
	helper.WriteResponBody(writer, webRespone)
}
