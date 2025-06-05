package eventcontroller

import (
	"net/http"
	"strconv"

	"backend/helper"
	"backend/model/dto"
	"backend/model/dto/eventdto"
	"backend/service/interfaces"

	"github.com/julienschmidt/httprouter"
)

type EventControllerImpl struct {
	EventService interfaces.EventService
}

func NewEventController(eventService interfaces.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (controller *EventControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Parse multipart form with 10MB max memory
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	eventCreateRequest := eventdto.EventCreateRequest{
		Heading:    request.FormValue("heading"),
		Subheading: request.FormValue("subheading"),
	}

	// Get the file from form data
	file, fileHeader, err := request.FormFile("thumbnail")
	helper.PanicIfError(err)
	defer file.Close()

	eventCreateRequest.Thumbnail = fileHeader

	eventResponse := controller.EventService.Create(request.Context(), eventCreateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Event created successfully",
		Data:    eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EventControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Parse multipart form with 10MB max memory
	err := request.ParseMultipartForm(10 << 20)
	helper.PanicIfError(err)

	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	eventUpdateRequest := eventdto.EventUpdateRequest{
		Id:         uint(id),
		Heading:    request.FormValue("heading"),
		Subheading: request.FormValue("subheading"),
	}

	// Get the file from form data
	file, fileHeader, err := request.FormFile("thumbnail")
	helper.PanicIfError(err)
	defer file.Close()

	eventUpdateRequest.Thumbnail = fileHeader

	eventResponse := controller.EventService.Update(request.Context(), eventUpdateRequest)
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Event updated successfully",
		Data:    eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EventControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	controller.EventService.Delete(request.Context(), uint(id))
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Event deleted successfully",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EventControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventId := params.ByName("eventId")
	id, err := strconv.Atoi(eventId)
	helper.PanicIfError(err)

	eventResponse := controller.EventService.FindById(request.Context(), uint(id))
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Event found successfully",
		Data:    eventResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EventControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	eventResponses := controller.EventService.FindAll(request.Context())
	webResponse := dto.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Events found successfully",
		Data:    eventResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
