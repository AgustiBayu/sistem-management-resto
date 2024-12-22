package main

import (
	"SistemManagementResto/app"
	"SistemManagementResto/controller"
	"SistemManagementResto/helper"
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
	router.POST("/api/menuItems", menuItemController.Create)
	router.GET("/api/menuItems", menuItemController.FindAll)
	router.GET("/api/menuItems/:menuItemId", menuItemController.FindById)
	router.PUT("/api/menuItems/:menuItemId", menuItemController.Update)
	router.DELETE("/api/menuItems/:menuItemId", menuItemController.Delete)
	router.POST("/api/pesanans", pesananController.Create)
	router.GET("/api/pesanans", pesananController.FindAll)
	router.GET("/api/pesanans/:pesananId", pesananController.FindById)
	router.PUT("/api/pesanans/:pesananId", pesananController.Update)
	router.DELETE("/api/pesanans/:pesananId", pesananController.Delete)
	router.POST("/api/detailPesanans", detailPesananController.Create)
	router.GET("/api/detailPesanans", detailPesananController.FindAll)
	router.GET("/api/detailPesanans/:detailPesananId", detailPesananController.FindById)
	router.PUT("/api/detailPesanans/:detailPesananId", detailPesananController.Update)
	router.DELETE("/api/detailPesanans/:detailPesananId", detailPesananController.Delete)
	router.POST("/api/transaksis", transaksiController.Create)
	router.GET("/api/transaksis", transaksiController.FindAll)
	router.GET("/api/transaksis/:transaksiId", transaksiController.FindById)
	router.PUT("/api/transaksis/:transaksiId", transaksiController.Update)
	router.DELETE("/api/transaksis/:transaksiId", transaksiController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
