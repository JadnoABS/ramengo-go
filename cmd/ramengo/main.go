package main

import (
 "log"
 "net/http"
 "github.com/joho/godotenv"
 "ramengo-go/internal/handlers"
 "ramengo-go/internal/middleware"
)

func main() {
	godotenv.Load()

	middleware := middleware.NewApiKeyMiddleware()

	http.Handle("/broths", middleware.Apply(handlers.GetBroths))
	http.Handle("/proteins", middleware.Apply(handlers.GetProteins))
	http.Handle("/orders", middleware.Apply(handlers.PostOrder))
	log.Fatal(http.ListenAndServe(":8080", nil))
}