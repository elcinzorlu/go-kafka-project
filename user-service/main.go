package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Gorilla Mux router'ını oluşturun.
	router := mux.NewRouter()

	// SetUserRoutes fonksiyonunu kullanarak rotaları ayarlayın.
	SetUserRoutes(router)

	// HTTP sunucusunu başlatın.
	port := 8080
	fmt.Printf("User Service is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
