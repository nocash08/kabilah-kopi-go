package userscontroller

import (
	"backend/helper"
	"backend/model/dto"
	"backend/model/dto/usersdto"
	"backend/service/interfaces"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersControllerImpl struct {
	UsersService interfaces.UsersService
}

func NewUsersController(usersService interfaces.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: usersService,
	}
}

func (controller *UsersControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersLoginRequest := usersdto.UsersLoginRequest{}

	helper.ReadFromRequestBody(request, &usersLoginRequest)

	_, tokenResponse := controller.UsersService.Login(request.Context(), writer, usersLoginRequest)

	// Create response with token
	response := usersdto.UsersResponse{
		Token: tokenResponse.AccessToken,
	}

	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful",
		Data:    response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := controller.UsersService.Logout(request.Context(), writer, request)
	if err != nil {
		webResponse := dto.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Logout successful",
	}

	helper.WriteToResponseBody(writer, webResponse)

	helper.ClearTokenCookie(writer, helper.AccessTokenCookie)
}

func (controller *UsersControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersRegisterRequest := usersdto.UsersRegisterRequest{}
	helper.ReadFromRequestBody(request, &usersRegisterRequest)

	user := controller.UsersService.Register(request.Context(), usersRegisterRequest)

	webResponse := dto.WebResponse{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "User registered successfully",
		Data:    user,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
