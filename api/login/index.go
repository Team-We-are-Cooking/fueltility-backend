package login

import (
	"encoding/json"
	"net/http"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	fueltilityhttp "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/http"
	fueltilitysupabase "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/supabase"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	crw := &fueltilityhttp.ResponseWriter{W: w}
	crw.SetCors(r.Host)

	method := r.Method

	_, err := fueltilitysupabase.CreateClient()
	if err != nil {
		crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
			Success: false,
			Error: &fueltilityhttp.ErrorDetails{Message: "Unable to connect to database."},
		})
	}

	switch method {
	case "GET":
		var data []schema.User = []schema.User{
			{Username: "bv", Password: "", Email: "email", FirstName: "Brandon", LastName: "C", Address: "address", City: "city", State: "TX", ZipCode: "134555"},
		}

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.User]{
			Success: true,
			Data: data,
		})
	case "POST":
		var userCreds schema.Credentials

		if err := json.NewDecoder(r.Body).Decode(&userCreds); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Internal server error."},
			})
		}

		var data []schema.Credentials
		
		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.Credentials]{
			Success: true,
			Data: data,
		})
	}
}