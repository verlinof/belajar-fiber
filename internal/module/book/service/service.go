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

func (b *BookService) CreateBook(ctx context.Context, createBookRequest book_model.CreateBookRequest) (book_model.BookResponse, error) {
	var book *book_model.Book
	var bookResponse book_model.BookResponse

	book = &book_model.Book{
		Title:       createBookRequest.Title,
		Description: createBookRequest.Description,
	}

	err := db.DB.WithContext(ctx).Table("books").Create(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	bookResponse = book_model.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
	}

	return bookResponse, nil
}

func (b *BookService) UpdateBook(ctx context.Context, id int, updateBookRequest book_model.UpdateBookRequest) (book_model.BookResponse, error) {
	var book book_model.Book
	err := db.DB.WithContext(ctx).Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	book.Title = updateBookRequest.Title
	book.Description = updateBookRequest.Description

	err = db.DB.WithContext(ctx).Table("books").Save(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	bookResponse := book_model.BookResponse(book)

	return bookResponse, nil
}

func (b *BookService) DeleteBook(ctx context.Context, id int) error {
	var book book_model.Book
	err := db.DB.WithContext(ctx).Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		return err
	}

	err = db.DB.WithContext(ctx).Table("books").Where("id = ?", id).Delete(&book_model.Book{}).Error
	if err != nil {
		return err
	}

	return nil
}
