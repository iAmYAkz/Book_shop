package routes

import (
	"yakz/handlers"
	"yakz/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserApi(app *fiber.App) {
	userGroup := app.Group("user",middleware.JWTMiddleware)

	userGroup.Delete("/:id",middleware.AdminMiddleware,handlers.DeleteUser)
	userGroup.Put("/:id/role",middleware.AdminMiddleware,handlers.UpdateUserRole)
	userGroup.Put("/:id/password",handlers.ChangeUserPassword)
	userGroup.Put("/:id",handlers.UpdateProfile)
	
	app.Post("/register", handlers.RegisterNewUser)
	app.Post("/Login", handlers.LoginUserNew)


}

func SetupBookApi(app *fiber.App) {
	bookGroup := app.Group("/books")

	bookGroup.Get("/", handlers.GetBooks)
	bookGroup.Get("/:id", handlers.GetBook)
	bookGroup.Post("/", middleware.JWTMiddleware, middleware.AdminMiddleware, handlers.CreateBook)
	bookGroup.Put("/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, handlers.UpdateBook)
	bookGroup.Delete("/:id", middleware.JWTMiddleware, middleware.AdminMiddleware, handlers.DeleteBook)

}

func SetupCartApi(app *fiber.App) {
	cartGroup := app.Group("/cart", middleware.JWTMiddleware)

	cartGroup.Post("/", handlers.AddtoCart)
	cartGroup.Get("/", handlers.GetCart)
	cartGroup.Put("/:book_id", handlers.UpdatrCartQty)
	cartGroup.Get("/Total", handlers.GetCartTotal)
	cartGroup.Delete("/clear", handlers.ClearCart)
	cartGroup.Delete("/:book_id", handlers.DeleteCart)
}
