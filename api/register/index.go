package register

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Team-We-are-Cooking/fueltility-backend/encrypt"
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
		log.Println(err.Error())
		crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
			Success: false,
			Error: &fueltilityhttp.ErrorDetails{Message: "Unable to connect to database."},
		})
	}

	switch method {
	case "POST":
		var userCreds schema.AuthCredentials

		if err := json.NewDecoder(r.Body).Decode(&userCreds); err != nil {
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Internal server error."},
			})

			return
		}

		hashdPw, err := encrypt.HashPassword(userCreds.Password)

		if err != nil {
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Internal server error."},
			})

			return
		}

		userCreds.Password = string(hashdPw)

		var createdUser schema.User

		if _, err := client.From("User").Insert(&userCreds, false, "", "", "").Single().ExecuteTo(&createdUser); err != nil{
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "User already exists."},
			})

			return
		}

		var data []schema.ReturnedCredentials = make([]schema.ReturnedCredentials, 1)
		data[0] = schema.ReturnedCredentials{ID: createdUser.ID, Username: createdUser.Username,  Email: createdUser.Email}
		
		crw.SendJSONResponse(http.StatusAccepted, fueltilityhttp.Response[schema.ReturnedCredentials]{
			Success: true,
			Data: data,
		})
	}
}