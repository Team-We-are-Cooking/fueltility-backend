package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Team-We-are-Cooking/fueltility-backend/api/fuel_quote"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/login"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/pricing_module"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/profile"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/register"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	
	mux.Handle("/api/fuel_quote", http.HandlerFunc(fuel_quote.Handler))
	mux.Handle("/api/login", http.HandlerFunc(login.Handler))
	mux.Handle("/api/pricing_module", http.HandlerFunc(pricing_module.Handler))
	mux.Handle("/api/profile", http.HandlerFunc(profile.Handler))
	mux.Handle("/api/register", http.HandlerFunc(register.Handler))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        c.Handler(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server listening on Port 8080. Live at http://localhost:8080")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}