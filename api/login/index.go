package login

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	fueltilityhttp "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/http"
	fueltilitysupabase "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/supabase"
	"golang.org/x/crypto/bcrypt"
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

		return
	}

	switch method {
	case "POST":
		var userCreds schema.Credentials

		if err := json.NewDecoder(r.Body).Decode(&userCreds); err != nil {
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Internal server error."},
			})

			return
		}
		
		var foundUser schema.Credentials
		if _, err := client.From("User").Select("*", "", false).Eq("username", userCreds.Username).Single().ExecuteTo(&foundUser); err != nil {
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusUnauthorized, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Invalid username or password."},
			})

			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userCreds.Password)); err != nil {
			log.Println(err.Error())
			crw.SendJSONResponse(http.StatusUnauthorized, fueltilityhttp.ErrorResponse{
				Success: false,
				Error: &fueltilityhttp.ErrorDetails{Message: "Invalid username or password."},
			})

			return
		}

		var data []schema.Credentials = make([]schema.Credentials, 1)
		data[0] = schema.Credentials{Username: foundUser.Username,  Email: foundUser.Email}
		
		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.Credentials]{
			Success: true,
			Data: data,
		})
	}
}