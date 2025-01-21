package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/configs/redis_config"
	book_http "github.com/verlinof/fiber-project-structure/internal/module/book/http"
	book_http_route "github.com/verlinof/fiber-project-structure/internal/module/book/http/route"
	book_service "github.com/verlinof/fiber-project-structure/internal/module/book/service"
	pkg_redis "github.com/verlinof/fiber-project-structure/pkg/redis"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

func InitRoute(app *fiber.App) {
	api := app.Group("/api")

	//Dependencies
	redisManager := pkg_redis.NewRedisManager(redis_config.Config.Host, redis_config.Config.Password, redis_config.Config.Db)
	validator := pkg_validation.NewXValidator()

	// Books
	bookService := book_service.NewBookService()
	bookHandler := book_http.NewBookHandler(bookService, redisManager, validator)
	book_http_route.BookRoute(api, bookHandler)
}
