package login

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	"github.com/joho/godotenv"
)

func Test_LoginHandler(t *testing.T) {
	data := []struct {
		Username		   string
		Email              string
		Password           string
		ExpectedStatusCode int
	}{
		{Username: "", Email: "", Password: "", ExpectedStatusCode: 400},
		{Username: "testaccount", Email: "", Password: "pw", ExpectedStatusCode: 400},
		{Username: "testaccount", Email: "test@gmail.com", Password: "", ExpectedStatusCode: 400},
		{Username: "notindb", Email: "notindb@yahoo.com", Password: "pw", ExpectedStatusCode: 401},
		{Username: "testaccount", Email: "asff@gmail.com", Password: "pw", ExpectedStatusCode: 401},
		{Username: "testaccount", Email: "test@gmail.com", Password: "123", ExpectedStatusCode: 200},
	}
	
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Unable to load environment variables: %s", err.Error())
	}

	for _, d := range data {
		t.Run("Test login api route", func(t *testing.T) {
			userCreds := schema.AuthCredentials{Username: d.Username, Email: d.Email, Password: d.Password}

			jsonData, err := json.Marshal(userCreds)
			if err != nil {
				t.Fatal(err)
			}

			r := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
			w := httptest.NewRecorder()
			handler := http.Handler(http.HandlerFunc(Handler))

			r.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(w, r)

			if status := w.Code; status != d.ExpectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, d.ExpectedStatusCode)
			}
		})
	}

	t.Run("Test error loading database", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/api/login", nil)
		w := httptest.NewRecorder()
		handler := http.Handler(http.HandlerFunc(Handler))

		handler.ServeHTTP(w, r)

		if status := w.Code; status != 500 {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, 500)
		}
	})
}