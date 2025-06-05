package implementation

import (
	"backend/exception"
	"backend/helper"
	"backend/helper/mapper"
	"backend/model/domain"
	"backend/model/dto/menudto"
	"backend/repository/menurepository"
	"backend/service/interfaces"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

type MenuServiceImpl struct {
	MenuRepository menurepository.MenuRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewMenuService(menuRepository menurepository.MenuRepository, db *sql.DB, validator *validator.Validate) interfaces.MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		DB:             db,
		Validator:      validator,
	}
}

func (service *MenuServiceImpl) Create(ctx context.Context, request menudto.MenuCreateRequest) menudto.MenuResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	uploadDir := filepath.Join("uploads", "thumbnails")
	thumbnailPath, err := helper.UploadFile(request.Thumbnail, uploadDir)
	helper.PanicIfError(err)

	menu := domain.Menu{
		Heading:    request.Heading,
		Subheading: request.Subheading,
		Thumbnail:  thumbnailPath,
	}

	menu = service.MenuRepository.Create(ctx, tx, menu)

	return mapper.ToMenuResponse(menu)
}

func (service *MenuServiceImpl) Update(ctx context.Context, request menudto.MenuUpdateRequest) menudto.MenuResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu, err := service.MenuRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if menu.Thumbnail != "" {
		if err := helper.DeleteFile(menu.Thumbnail); err != nil {
			fmt.Printf("Warning: failed to delete old thumbnail: %v\n", err)
		}
	}

	uploadDir := filepath.Join("uploads", "thumbnails")
	thumbnailPath, err := helper.UploadFile(request.Thumbnail, uploadDir)
	helper.PanicIfError(err)

	menu.Heading = request.Heading
	menu.Subheading = request.Subheading
	menu.Thumbnail = thumbnailPath

	menu = service.MenuRepository.Update(ctx, tx, menu)

	return mapper.ToMenuResponse(menu)
}

func (service *MenuServiceImpl) Delete(ctx context.Context, menuId uint) {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu, err := service.MenuRepository.FindById(ctx, tx, menuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MenuRepository.Delete(ctx, tx, menu)
}

func (service *MenuServiceImpl) FindAll(ctx context.Context) []menudto.MenuResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menus := service.MenuRepository.FindAll(ctx, tx)

	return mapper.ToMenuResponses(menus)
}

func (service *MenuServiceImpl) FindById(ctx context.Context, menuId uint) menudto.MenuResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	menu, err := service.MenuRepository.FindById(ctx, tx, menuId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return mapper.ToMenuResponse(menu)
}
