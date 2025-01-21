package book_http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	book_model "github.com/verlinof/fiber-project-structure/internal/module/book/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	"gorm.io/gorm"
)

func (bookHandler *BookHandler) GetAllBook(c *fiber.Ctx) error {
	//Pagination
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per_page", 10)

	books, err := bookHandler.bookService.GetAllBook(c.Context(), page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (bookHandler *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// Jika konversi gagal, kembalikan error ke klien
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Error Handling
	book, err := bookHandler.bookService.GetBookByID(c.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (bookHandler *BookHandler) CreateBook(c *fiber.Ctx) error {
	var createBookRequest book_model.CreateBookRequest
	err := c.BodyParser(&createBookRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Validate the struct
	err = bookHandler.xValidator.Validate(createBookRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Error Handling
	book, err := bookHandler.bookService.CreateBook(c.Context(), createBookRequest)
	if err != nil {
		//Err Duplicated Unique Key
		if err == gorm.ErrDuplicatedKey {
			return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (bookHandler *BookHandler) UpdateBook(c *fiber.Ctx) error {
	var updateBookRequest book_model.UpdateBookRequest
	err := c.BodyParser(&updateBookRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// Validate struct
	err = bookHandler.xValidator.Validate(updateBookRequest)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// error handling
	book, err := bookHandler.bookService.UpdateBook(c.Context(), id, updateBookRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		} else if err == gorm.ErrDuplicatedKey {
			return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (bookHandler *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// Jika konversi gagal, kembalikan error ke klien
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	err = bookHandler.bookService.DeleteBook(c.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return c.Status(http.StatusNoContent).Send([]byte{})
}
