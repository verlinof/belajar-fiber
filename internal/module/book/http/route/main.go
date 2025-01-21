package book_http_route

import (
	"github.com/gofiber/fiber/v2"
	book_http "github.com/verlinof/fiber-project-structure/internal/module/book/http"
)

func BookRoute(router fiber.Router, bookHandler book_http.BookHandler) {
	bookRoutes := router.Group("/books")
	bookRoutes.Get("/", bookHandler.GetAllBook)
	bookRoutes.Get("/:id", bookHandler.GetBookByID)
	bookRoutes.Post("/", bookHandler.CreateBook)
	bookRoutes.Patch("/:id", bookHandler.UpdateBook)
	bookRoutes.Delete("/:id", bookHandler.DeleteBook)
}
