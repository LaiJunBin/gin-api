package main

import (
	"github.com/LaiJunBin/gin-api/internal/routers"
	"net/http"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
	}

	s.ListenAndServe()
}
