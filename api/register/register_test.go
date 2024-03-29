package register

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	"github.com/joho/godotenv"
)

func Test_RegisterHandler(t *testing.T) {
	data := []struct {
		Username           string
		Email              string
		Password           string
		ExpectedStatusCode int
	}{
		{Username: "", Email: "", Password: "", ExpectedStatusCode: 400},
		{Username: "testaccount", Email: "", Password: "pw", ExpectedStatusCode: 400},
		{Username: "testaccount", Email: "test@gmail.com", Password: "", ExpectedStatusCode: 400},
		{Username: "testaccount", Email: "test@gmail.com", Password: "123", ExpectedStatusCode: 409},
		{Username: "testaccount2", Email: "test2@gmail.com", Password: "123", ExpectedStatusCode: 409},
	}

	for _, d := range data {
		t.Run("Test register api route", func(t *testing.T) {
			if err := godotenv.Load("../../.env"); err != nil {
				t.Fatalf("Unable to load environment variables: %s", err.Error())
			}

			userCreds := schema.AuthCredentials{Username: d.Username, Email: d.Email, Password: d.Password}

			jsonData, err := json.Marshal(userCreds)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()
			handler := http.Handler(http.HandlerFunc(Handler))

			r.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}

	t.Run("Test loading database error", func(t *testing.T) {
		t.Setenv("SUPABASE_URL", "")
		t.Setenv("SUPABASE_KEY", "")

		r := httptest.NewRequest("POST", "/api/register", nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusInternalServerError {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
		}
	})

	t.Run("Test no http request body error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/api/register", nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusBadRequest {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}
