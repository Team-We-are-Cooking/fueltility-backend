package pricing_module

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	"github.com/joho/godotenv"
)

func Test_Pricing_ModuleHandler(t *testing.T) {
	t.Run("Test pricing_module api route", func(t *testing.T) {
		// Load environment variables from .env file
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load environment variables: %s", err.Error())
		}
		 // Create a request to pass to our handler.
		req, err := http.NewRequest("GET", "/api/pricing_module", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}



	})

	t.Run("Test error loading database", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/api/pricing_module", nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != 500 {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, 500)
		}
	})

}