package book_http

import (
	"github.com/gofiber/fiber/v2"
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
