package exception

import (
	"backend/helper"
	"backend/model/dto"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := dto.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Message: exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := dto.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT FOUND",
			Message: exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// Convert error to string for detailed message
	var errorMessage string
	if err != nil {
		switch v := err.(type) {
		case error:
			errorMessage = v.Error()
		case string:
			errorMessage = v
		default:
			errorMessage = fmt.Sprintf("%v", v)
		}
	} else {
		errorMessage = "Unknown error occurred"
	}

	webResponse := dto.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL SERVER ERROR",
		Message: errorMessage,
	}

	// Log the error
	fmt.Printf("Internal Server Error: %v\n", errorMessage)

	helper.WriteToResponseBody(w, webResponse)
}
