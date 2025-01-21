package book_service

import (
	"context"
	"math"

	"github.com/verlinof/fiber-project-structure/db"
	book_model "github.com/verlinof/fiber-project-structure/internal/module/book/model"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
)

func (b *BookService) GetAllBook(ctx context.Context, page int, perPage int) (*pkg_success.PaginationData, error) {
	var books []book_model.BookResponse
	var totalRows int64

	//Pagination System
	offset := (page - 1) * perPage
	db.DB.WithContext(ctx).Table("books").Count(&totalRows)

	totalPage := math.Ceil(float64(totalRows) / float64(perPage))

	err := db.DB.WithContext(ctx).Limit(perPage).Offset(offset).Table("books").Find(&books).Error
	if err != nil {
		return nil, err
	}

	//Response
	response := pkg_success.SuccessPaginationData(books, page, int(totalPage), perPage, int(totalRows))

	return response, nil
}

func (b *BookService) GetBookByID(ctx context.Context, id int) (book_model.BookResponse, error) {
	var book book_model.BookResponse
	err := db.DB.WithContext(ctx).Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	return book, nil
}
