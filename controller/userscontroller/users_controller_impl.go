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

	usersResponse, tokenResponse := controller.UsersService.Login(request.Context(), writer, usersLoginRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful",
		Data:    usersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

	helper.SetTokenCookie(writer, helper.AccessTokenCookie, tokenResponse.AccessToken, helper.AccessTokenDuration)

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) RefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tokenResponse, err := controller.UsersService.RefreshToken(request.Context(), writer, request)
	helper.PanicIfError(err)

	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Refresh token successful",
		Data:    tokenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

	helper.SetTokenCookie(writer, helper.AccessTokenCookie, tokenResponse.AccessToken, helper.AccessTokenDuration)
}

func (controller *UsersControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := controller.UsersService.Logout(request.Context(), writer)
	helper.PanicIfError(err)

	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Logout successful",
	}

	helper.WriteToResponseBody(writer, webResponse)

	helper.ClearTokenCookie(writer, helper.AccessTokenCookie)
	helper.ClearTokenCookie(writer, helper.RefreshTokenCookie)
}
