package fuel_quote

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

	quote_id := r.URL.Query().Get("quote_id")
	user_id := r.URL.Query().Get("user_id")

	if user_id == "" && quote_id == "" {
		switch method {
		case "POST":
			var userFuelQuoteData schema.FuelQuote

			if err := json.NewDecoder(r.Body).Decode(&userFuelQuoteData); err != nil {
				crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
					Success: false,
					Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
				})
				return
			}
			
			temp, _, err := client.From("Fuel Quote").Insert(userFuelQuoteData, false, "", "", "exact").Execute(); 
			if err != nil {
				crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
					Success: false,
					Error:   &fueltilityhttp.ErrorDetails{Message: "Failed to create user: " + err.Error()},
				})
				return
			}

			var fuelQuote []schema.FuelQuote
			err = json.Unmarshal(temp, &fuelQuote)
			if err != nil {
				crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
					Success: false,
					Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
				})
				return
			}

			crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.FuelQuote]{
				Success: true,
				Data: fuelQuote,
			})			
		default:
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "No other method allowed without any parameters."},
			})
		}
	} else if user_id != "" {
		switch method {
		case "GET":
			var data []schema.FuelQuote
			if _, err := client.From("Fuel Quote").Select("*", "exact", false).Eq("user_id", user_id).ExecuteTo(&data); err != nil {
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
		default:
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "No other method allowed with user_id."},
			})
		}
	} else if quote_id != "" {
		switch method {
		case "GET":
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
		default:
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "No other method allowed with quote_id."},
			})
		}
	}
}
