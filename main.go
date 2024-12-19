package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Swagger UI handler
	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.StripPrefix("/swagger", http.FileServer(http.Dir("./openapi_docs"))).ServeHTTP(w, r)
	})

	// Start server
	log.Println("Server is running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
