package menucontroller

import (
	"backend/model/dto/menudto"
	"backend/service/interfaces"
	"net/http"
	"strconv"

	"backend/helper"
	"backend/model/dto"

	"github.com/julienschmidt/httprouter"
)

type MenuControllerImpl struct {
	MenuService interfaces.MenuService
}

func NewMenuController(menuService interfaces.MenuService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService,
	}
}

func (controller *MenuControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Parse multipart form with 10MB max memory
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	menuCreateRequest := menudto.MenuCreateRequest{
		Heading:    request.FormValue("heading"),
		Subheading: request.FormValue("subheading"),
	}

	// Get the file from form data
	file, fileHeader, err := request.FormFile("thumbnail")
	helper.PanicIfError(err)
	defer file.Close()

	menuCreateRequest.Thumbnail = fileHeader

	menuResponse := controller.MenuService.Create(request.Context(), menuCreateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Menu created successfully",
		Data:    menuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Parse multipart form with 10MB max memory
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	menuId := params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
	helper.PanicIfError(err)

	menuUpdateRequest := menudto.MenuUpdateRequest{
		Id:         uint(id),
		Heading:    request.FormValue("heading"),
		Subheading: request.FormValue("subheading"),
	}

	// Get the file from form data
	file, fileHeader, err := request.FormFile("thumbnail")
	helper.PanicIfError(err)
	defer file.Close()

	menuUpdateRequest.Thumbnail = fileHeader

	menuResponse := controller.MenuService.Update(request.Context(), menuUpdateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Menu updated successfully",
		Data:    menuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.MenuService.Delete(request.Context(), uint(id))
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Menu deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	menuResponse := controller.MenuService.FindById(request.Context(), uint(id))
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Menu found successfully",
		Data:    menuResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MenuControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	menuResponses := controller.MenuService.FindAll(request.Context())
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Menus found successfully",
		Data:    menuResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
