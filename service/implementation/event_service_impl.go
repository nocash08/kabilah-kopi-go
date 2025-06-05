package implementation

import (
	"backend/helper"
	"backend/helper/mapper"
	"backend/model/domain"
	"backend/model/dto/eventdto"
	"backend/repository/eventrepository"
	"backend/service/interfaces"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

type EventServiceImpl struct {
	EventRepository eventrepository.EventRepository
	DB              *sql.DB
	Validator       *validator.Validate
}

func NewEventService(eventRepository eventrepository.EventRepository, db *sql.DB, validator *validator.Validate) interfaces.EventService {
	return &EventServiceImpl{
		EventRepository: eventRepository,
		DB:              db,
		Validator:       validator,
	}
}

func (service *EventServiceImpl) Create(ctx context.Context, request eventdto.EventCreateRequest) eventdto.EventResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	uploadDir := filepath.Join("uploads", "thumbnails")
	thumbnailPath, err := helper.UploadFile(request.Thumbnail, uploadDir)
	helper.PanicIfError(err)

	event := domain.Event{
		Heading:    request.Heading,
		Subheading: request.Subheading,
		Thumbnail:  thumbnailPath,
	}

	event = service.EventRepository.Create(ctx, tx, event)

	return mapper.ToEventResponse(event)
}

func (service *EventServiceImpl) Update(ctx context.Context, request eventdto.EventUpdateRequest) eventdto.EventResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.EventRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	if event.Thumbnail != "" {
		if err := helper.DeleteFile(event.Thumbnail); err != nil {
			fmt.Printf("Warning: failed to delete old thumbnail: %v\n", err)
		}
	}

	uploadDir := filepath.Join("uploads", "thumbnails")
	thumbnailPath, err := helper.UploadFile(request.Thumbnail, uploadDir)
	helper.PanicIfError(err)

	event.Heading = request.Heading
	event.Subheading = request.Subheading
	event.Thumbnail = thumbnailPath

	event = service.EventRepository.Update(ctx, tx, event)

	return mapper.ToEventResponse(event)
}

func (service *EventServiceImpl) Delete(ctx context.Context, eventId uint) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.EventRepository.FindById(ctx, tx, eventId)
	helper.PanicIfError(err)

	service.EventRepository.Delete(ctx, tx, event)
}

func (service *EventServiceImpl) FindAll(ctx context.Context) []eventdto.EventResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	events := service.EventRepository.FindAll(ctx, tx)

	return mapper.ToEventResponses(events)
}

func (service *EventServiceImpl) FindById(ctx context.Context, eventId uint) eventdto.EventResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	event, err := service.EventRepository.FindById(ctx, tx, eventId)
	helper.PanicIfError(err)

	return mapper.ToEventResponse(event)
}
