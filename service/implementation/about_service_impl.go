package implementation

import (
	"context"
	"database/sql"

	"backend/helper"
	"backend/helper/mapper"
	"backend/model/domain"
	"backend/model/dto/aboutdto"
	"backend/repository/aboutrepository"
	"backend/service/interfaces"

	"github.com/go-playground/validator/v10"
)

type AboutServiceImpl struct {
	AboutRepository aboutrepository.AboutRepository
	DB              *sql.DB
	Validator       *validator.Validate
}

func NewAboutService(aboutRepository aboutrepository.AboutRepository, db *sql.DB, validator *validator.Validate) interfaces.AboutService {
	return &AboutServiceImpl{
		AboutRepository: aboutRepository,
		DB:              db,
		Validator:       validator,
	}
}

func (service *AboutServiceImpl) Create(ctx context.Context, request aboutdto.AboutCreateRequest) aboutdto.AboutResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	about := domain.About{
		Heading:    request.Heading,
		Subheading: request.Subheading,
	}

	about = service.AboutRepository.Create(ctx, tx, about)

	return mapper.ToAboutResponse(about)
}

func (service *AboutServiceImpl) Update(ctx context.Context, request aboutdto.AboutUpdateRequest) aboutdto.AboutResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	about, err := service.AboutRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	about.Heading = request.Heading
	about.Subheading = request.Subheading

	about = service.AboutRepository.Update(ctx, tx, about)

	return mapper.ToAboutResponse(about)
}

func (service *AboutServiceImpl) FindAll(ctx context.Context) []aboutdto.AboutResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	abouts := service.AboutRepository.FindAll(ctx, tx)

	return mapper.ToAboutResponses(abouts)
}

func (service *AboutServiceImpl) FindById(ctx context.Context, aboutId uint) aboutdto.AboutResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	about, err := service.AboutRepository.FindById(ctx, tx, aboutId)
	helper.PanicIfError(err)

	return mapper.ToAboutResponse(about)
}
