package main

import (
	"fmt"
	"net/http"
	"silverApi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Silver API")
	err := http.ListenAndServe("localhost:8082", r)
	if err != nil {
		log.Error(err)
	}
}
