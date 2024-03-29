package profile

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	"github.com/joho/godotenv"
)

func Test_ProfileHandler(t *testing.T) {
	GETReqData := []struct {
		UserID string
		ExpectedStatusCode int
	}{
		{UserID: "", ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec619", ExpectedStatusCode: http.StatusOK},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec664", ExpectedStatusCode: http.StatusOK},
	}

	PUTReqData := []struct {
		UserID             string
		RequestBody        interface{}
		ExpectedStatusCode int
	}{
		{UserID: "", RequestBody: schema.Profile{}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "", RequestBody: nil, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "", LastName: "", Address: "123address", AddressTwo: "address2", City: "Houston", State: "TX", ZipCode: "12345"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "", AddressTwo: "address2", City: "Houston", State: "TX", ZipCode: "12345"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "address1", AddressTwo: "address2address2address2address2address2address2address2address2address2address2address2address2address2address2address2", City: "Houston", State: "TX", ZipCode: "12345"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "123address", AddressTwo: "address2", City: "", State: "TX", ZipCode: "12345"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "123address", AddressTwo: "address2", City: "Houston", State: "", ZipCode: "12345"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0490a9ec623", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "123address", AddressTwo: "address2", City: "Houston", State: "", ZipCode: "12"}, ExpectedStatusCode: http.StatusBadRequest},
		{UserID: "70916454-df22-4b07-882e-f0ddddddddddddddddddddddddddddddd", RequestBody: schema.Profile{FirstName: "John", LastName: "Doe", Address: "123address", AddressTwo: "address2", City: "Houston", State: "TX", ZipCode: "12345"}, ExpectedStatusCode: http.StatusInternalServerError},
	}


	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Unable to load environment variables: %s", err.Error())
	}

	for _, d := range GETReqData {
		t.Run("Test fuel quote handler GET Request", func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api/profile?user_id="+d.UserID, nil)
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)
			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}

	for _, d := range PUTReqData {
		t.Run("Test fuel quote handler PUT Request", func(t *testing.T) {
			var jsonData []byte
			
			if d.RequestBody != nil {
				var err error
				jsonData, err = json.Marshal(d.RequestBody)
				if err != nil {
					t.Fatal(err)
				}
			}

			r := httptest.NewRequest("PUT", "/api/profile?user_id="+d.UserID, bytes.NewBuffer(jsonData))
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
