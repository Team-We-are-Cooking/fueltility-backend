package pricing_module

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func Test_Pricing_ModuleHandler(t *testing.T) {
	t.Run("Test pricing_module api route", func(t *testing.T) {
		// Load environment variables from .env file
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load environment variables: %s", err.Error())
		}

		quote_id := "1"

		req := httptest.NewRequest("GET", "/api/pricing_module?quote_id="+quote_id, nil)

		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})

	t.Run("Test pricing_module api route missing query error", func(t *testing.T) {
		// Load environment variables from .env file
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load environment variables: %s", err.Error())
		}

		quote_id := ""

		req := httptest.NewRequest("GET", "/api/pricing_module?quote_id="+quote_id, nil)

		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("Test error loading database", func(t *testing.T) {
		t.Setenv("SUPABASE_URL", "")
		t.Setenv("SUPABASE_KEY", "")

		quote_id := "1"

		r := httptest.NewRequest("GET", "/api/fuel_quote?quote_id="+quote_id, nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != 500 {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, 500)
		}
	})

}