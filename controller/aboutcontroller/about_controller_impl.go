package aboutcontroller

import (
	"net/http"
	"strconv"

	"backend/helper"
	"backend/model/dto"
	"backend/model/dto/aboutdto"
	"backend/service/interfaces"

	"github.com/julienschmidt/httprouter"
)

type AboutControllerImpl struct {
	AboutService interfaces.AboutService
}

func NewAboutController(aboutService interfaces.AboutService) AboutController {
	return &AboutControllerImpl{
		AboutService: aboutService,
	}
}

func (controller *AboutControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aboutCreateRequest := aboutdto.AboutCreateRequest{}

	helper.ReadFromRequestBody(request, &aboutCreateRequest)

	aboutResponse := controller.AboutService.Create(request.Context(), aboutCreateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "About created successfully",
		Data:    aboutResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AboutControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aboutUpdateRequest := aboutdto.AboutUpdateRequest{}

	helper.ReadFromRequestBody(request, &aboutUpdateRequest)

	aboutId := params.ByName("aboutId")
	id, err := strconv.Atoi(aboutId)
	helper.PanicIfError(err)

	aboutUpdateRequest.Id = uint(id)

	aboutResponse := controller.AboutService.Update(request.Context(), aboutUpdateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "About updated successfully",
		Data:    aboutResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AboutControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aboutId := params.ByName("aboutId")
	id, err := strconv.Atoi(aboutId)
	helper.PanicIfError(err)

	aboutResponse := controller.AboutService.FindById(request.Context(), uint(id))
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "About found successfully",
		Data:    aboutResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AboutControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	aboutResponses := controller.AboutService.FindAll(request.Context())
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Abouts found successfully",
		Data:    aboutResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
