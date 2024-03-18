package fuel_quote

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	fueltilityhttp "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/http"
	fueltilitysupabase "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/supabase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	crw := &fueltilityhttp.ResponseWriter{W: w}
	crw.SetCors(r.Host)

	method := r.Method

	client, err := fueltilitysupabase.CreateClient()
	if err != nil {
		crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
			Success: false,
			Error:   &fueltilityhttp.ErrorDetails{Message: "Unable to connect to database."},
		})
		return
	}

	switch method {
	case "GET":
		quote_id := r.URL.Query().Get("profile")
		if quote_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing quote id."},
			})
			return
		}

		var data []schema.FuelQuote
		if _, err := client.From("Fuel Quote").Select("*", "exact", false).Eq("quote_id", quote_id).ExecuteTo(&data); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
			})
			return
		}

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.FuelQuote]{
			Success: true,
			Data:    data,
		})
	case "PUT":
		profile_id := r.URL.Query().Get("profile_id")
		if profile_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing profile id."},
			})
			return
		}

		var userProfile schema.User

		if err := json.NewDecoder(r.Body).Decode(&userProfile); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
			})
			return
		}

		if _, _, err := client.From("User").Update(userProfile, "", "exact").Eq("id", profile_id).Execute(); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Failed to update user: " + err.Error()},
			})
			return
		}

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.User]{
			Success: true,
		})
	}
}
