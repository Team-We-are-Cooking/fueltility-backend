package profile

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
		user_id := r.URL.Query().Get("user_id")
		if user_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing user id."},
			})
			return
		}

		var data []schema.User
		if _, err := client.From("User").Select("*", "exact", false).Eq("id", user_id).ExecuteTo(&data); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Failed to retrieve user: " + err.Error()},
			})
			return
		}

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.User]{
			Success: true,
			Data:    data,
		})
	case "PUT":
		user_id := r.URL.Query().Get("user_id")
		if user_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing profile id."},
			})
			return
		}

		var userProfile schema.Profile
		if err := json.NewDecoder(r.Body).Decode(&userProfile); err != nil {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
			})
			return
		}

		var updatedUser schema.User
		if _, err := client.From("User").Update(userProfile, "", "exact").Eq("id", user_id).Single().ExecuteTo(&updatedUser); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Failed to update user: " + err.Error()},
			})
			return
		}

		var data []schema.User = make([]schema.User, 1)
		data[0] = updatedUser

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.User]{
			Success: true,
			Data:    data,
		})
	}
}
