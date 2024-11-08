package exception

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

func ErrorHandler(rw http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(rw, r, err) {
		return
	}

	if validationError(rw, r, err) {
		return
	}
	internalServerError(rw, r, err)
}

func validationError(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		rw.Header().Set("Content-type", "application/json")
		rw.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(rw http.ResponseWriter, r *http.Request, err interface{}) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
