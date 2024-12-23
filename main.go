package main

import (
	"SistemManagementResto/app"
	"SistemManagementResto/controller"
	"SistemManagementResto/exception"
	"SistemManagementResto/helper"
	"SistemManagementResto/middleware"
	"SistemManagementResto/repository"
	"SistemManagementResto/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
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
	pesananRepository := repository.NewPesananRepository()
	pesananService := service.NewPesananService(pesananRepository, userRepository, db, validate)
	pesananController := controller.NewPesananController(pesananService)
	detailPesananRepository := repository.NewDetailPesananRepository()
	detailPesananService := service.NewDetailPesananService(detailPesananRepository, pesananRepository, userRepository, menuItemRepository, db, validate)
	detailPesananController := controller.NewDetailPesananController(detailPesananService)
	transaksiRepository := repository.NewTransaksiRepository()
	transaksiService := service.NewTransaksiService(transaksiRepository, pesananRepository, userRepository, db, validate)
	transaksiController := controller.NewTransaksiController(transaksiService)

	router := httprouter.New()
	// Swagger UI handler
	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.StripPrefix("/swagger", http.FileServer(http.Dir("./openapi_docs"))).ServeHTTP(w, r)
	})
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/menuItems", middleware.BasicAuth(userService, menuItemController.Create))
	router.GET("/api/menuItems", middleware.BasicAuth(userService, menuItemController.FindAll))
	router.GET("/api/menuItems/:menuItemId", middleware.BasicAuth(userService, menuItemController.FindById))
	router.PUT("/api/menuItems/:menuItemId", middleware.BasicAuth(userService, menuItemController.Update))
	router.DELETE("/api/menuItems/:menuItemId", middleware.BasicAuth(userService, menuItemController.Delete))
	router.POST("/api/pesanans", middleware.BasicAuth(userService, pesananController.Create))
	router.GET("/api/pesanans", middleware.BasicAuth(userService, pesananController.FindAll))
	router.GET("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.FindById))
	router.PUT("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.Update))
	router.DELETE("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.Delete))
	router.POST("/api/detailPesanans", middleware.BasicAuth(userService, detailPesananController.Create))
	router.GET("/api/detailPesanans", middleware.BasicAuth(userService, detailPesananController.FindAll))
	router.GET("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.FindById))
	router.PUT("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.Update))
	router.DELETE("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.Delete))
	router.POST("/api/transaksis", middleware.BasicAuth(userService, transaksiController.Create))
	router.GET("/api/transaksis", middleware.BasicAuth(userService, transaksiController.FindAll))
	router.GET("/api/transaksis/:transaksiId", middleware.BasicAuth(userService, transaksiController.FindById))
	router.PUT("/api/transaksis/:transaksiId", middleware.BasicAuth(userService, transaksiController.Update))
	router.DELETE("/api/transaksis/:transaksiId", middleware.BasicAuth(userService, transaksiController.Delete))

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
