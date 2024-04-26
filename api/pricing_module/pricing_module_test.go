package pricing_module

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func Test_Pricing_ModuleHandler(t *testing.T) {

	data := []struct {
		QuoteID            string
		Method             string
		ExpectedStatusCode int
	}{
		// Test cases for GET requests
		{QuoteID: "", Method: "GET", ExpectedStatusCode: http.StatusBadRequest},                 // missing quote id
		{QuoteID: "wrongID", Method: "GET", ExpectedStatusCode: http.StatusInternalServerError}, // wrong type of quote id
		{QuoteID: "4", Method: "GET", ExpectedStatusCode: http.StatusOK},                        // wrong type of quote id

	}

	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Unable to load environment variables: %s", err.Error())
	}

	for _, d := range data {
		t.Run("Test fuel quote handler", func(t *testing.T) {

			r := httptest.NewRequest(d.Method, "/api/fuel_quote?quote_id="+d.QuoteID, nil)
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)
			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}

	t.Run("Test error loading database", func(t *testing.T) {
		t.Setenv("SUPABASE_URL", "")
		t.Setenv("SUPABASE_KEY", "")

		r := httptest.NewRequest("POST", "/api/fuel_quote", nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != 500 {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, 500)
		}
	})
}
