package main

import (
	"SistemManagementResto/app"
	"SistemManagementResto/controller"
	"SistemManagementResto/helper"
	"SistemManagementResto/repository"
	"SistemManagementResto/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	menuItemRepository := repository.NewMenuItemRepository()
	menuItemService := service.NewMenuItemService(menuItemRepository, db, validate)
	menuItemController := controller.NewMenuItemController(menuItemService)

	router := httprouter.New()
	// Swagger UI handler
	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.StripPrefix("/swagger", http.FileServer(http.Dir("./openapi_docs"))).ServeHTTP(w, r)
	})
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/menuItems", menuItemController.Create)
	router.GET("/api/menuItems", menuItemController.FindAll)
	router.GET("/api/menuItems/:menuItemId", menuItemController.FindById)
	router.PUT("/api/menuItems/:menuItemId", menuItemController.Update)
	router.DELETE("/api/menuItems/:menuItemId", menuItemController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
