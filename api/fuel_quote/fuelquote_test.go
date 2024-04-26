package fuel_quote

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	"github.com/joho/godotenv"
)

func Test_FuelQuoteHandler(t *testing.T) {
	data := []struct {
		UserID             string
		QuoteID            string
		Method             string
		RequestBody        interface{}
		ExpectedStatusCode int
	}{
		// Test cases for POST requests
		{UserID: "", QuoteID: "", Method: "POST", RequestBody: nil, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "", QuoteID: "", Method: "POST", RequestBody: schema.FuelQuote{}, ExpectedStatusCode: http.StatusInternalServerError}, // Empty request body
		{UserID: "2d8d4210-0309-4940-9229-05a7a67a5d17", QuoteID: "", Method: "POST", RequestBody: schema.FuelQuote{}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "", QuoteID: "1", Method: "POST", RequestBody: schema.FuelQuote{}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "2d8d4210-0309-4940-9229-05a7a67a5d17", QuoteID: "1", Method: "POST", RequestBody: schema.FuelQuote{}, ExpectedStatusCode: http.StatusBadRequest},

		// Test cases for GET requests
		{UserID: "", QuoteID: "", Method: "GET", RequestBody: nil, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "2d8d4210-0309-4940-9229-05a7a67a5d17", QuoteID: "", Method: "GET", RequestBody: nil, ExpectedStatusCode: http.StatusOK},
		{UserID: "", QuoteID: "1", Method: "GET", RequestBody: nil, ExpectedStatusCode: http.StatusOK},
		{UserID: "wrongID", QuoteID: "", Method: "GET", RequestBody: nil, ExpectedStatusCode: http.StatusInternalServerError}, 
		{UserID: "", QuoteID: "wrongID", Method: "GET", RequestBody: nil, ExpectedStatusCode: http.StatusInternalServerError}, 

	}

	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Unable to load environment variables: %s", err.Error())
	}

	for _, d := range data {
		t.Run("Test fuel quote handler", func(t *testing.T) {
			var jsonData []byte
			if d.RequestBody != nil {
				var err error
				jsonData, err = json.Marshal(d.RequestBody)
				if err != nil {
					t.Fatal(err)
				}
			}

			r := httptest.NewRequest(d.Method, "/api/fuel_quote?quote_id="+d.QuoteID+"&user_id="+d.UserID, bytes.NewBuffer(jsonData))
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
