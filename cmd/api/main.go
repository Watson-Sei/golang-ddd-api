package main

import (
	"golang-ddd-api/app/config"
	"golang-ddd-api/app/infrastructure/persistence"
	"golang-ddd-api/app/interface/handler"
	"golang-ddd-api/app/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUserCase(userPersistence)
	userHandler := handler.NewUserhandler(userUseCase)

	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", userHandler.Index)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
