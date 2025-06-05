package userscontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	RefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
