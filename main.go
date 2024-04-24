package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Team-We-are-Cooking/fueltility-backend/api/fuel_quote"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/login"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/pricing_module"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/profile"
	"github.com/Team-We-are-Cooking/fueltility-backend/api/register"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	portNo := 8080

	mux := http.NewServeMux()
	
	mux.Handle("/api/fuel_quote", http.HandlerFunc(fuel_quote.Handler))
	mux.Handle("/api/login", http.HandlerFunc(login.Handler))
	mux.Handle("/api/pricing_module", http.HandlerFunc(pricing_module.Handler))
	mux.Handle("/api/profile", http.HandlerFunc(profile.Handler))
	mux.Handle("/api/register", http.HandlerFunc(register.Handler))

	log.Printf("Server live on port %d\n", portNo)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", portNo), mux); err != nil {
		log.Fatal(err)
	}
}