package pricing_module

import (
	"github.com/Team-We-are-Cooking/fueltility-backend/schema"
	fueltilityhttp "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/http"
	fueltilitysupabase "github.com/Team-We-are-Cooking/fueltility-backend/wrappers/supabase"
	"net/http"
	"strconv"
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
		quote_id := r.URL.Query().Get("quote_id")
		if quote_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing quote id."},
			})
			return
		}

		var data []schema.PricingModule
		if _, err := client.From("Pricing Module").Select("*", "exact", false).Eq("quote_id", quote_id).ExecuteTo(&data); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
			})
			return
		}

		// If no data is found, insert the event
		int_quote_id, _ := strconv.Atoi(quote_id)
		if len(data) == 0 {
			if _, _, err := client.From("Pricing Module").Insert(schema.PricingModule{QuoteId: int8(int_quote_id)}, false, "", "", "exact").Execute(); err != nil {
				crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
					Success: false,
					Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
				})
				return
			}

			if _, err := client.From("Pricing Module").Select("*", "exact", false).Eq("quote_id", quote_id).ExecuteTo(&data); err != nil {
				crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
					Success: false,
					Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
				})
				return
			}
		}

		//Fetch the FuelQuote based on quoteID
		var fuelQuote []schema.FuelQuote
		_, _ = client.From("Fuel Quote").Select("*", "exact", false).Eq("quote_id", quote_id).ExecuteTo(&fuelQuote)

		// Calculate the suggest price and total amount due for each fuel quote
		const pricePerGallon float32 = 1.50
		var margin float32
		var locationFactor float32
		var rateHistoryFactor float32
		var gallonsRequestedFactor float32
		var companyProfitFactor float32 = 0.1
		if fuelQuote[0].Interstate {
			locationFactor = 0.04
		} else {
			locationFactor = 0.02
		}
		if fuelQuote[0].GallonsRequested > 1000 {
			gallonsRequestedFactor = 0.02
		} else {
			gallonsRequestedFactor = 0.03
		}
		// Fetch the FuelQuote based on UserID
		var userFuelQuotes []schema.FuelQuote
		_, _ = client.From("Fuel Quote").Select("*", "exact", false).Eq("user_id", fuelQuote[0].UserId.String()).ExecuteTo(&userFuelQuotes)

		// Check if the user has requested fuel before
		if len(userFuelQuotes) > 0 {
			rateHistoryFactor = 0.01
		} else {
			rateHistoryFactor = 0.00
		}
		// Calculate the margin
		margin = pricePerGallon * (locationFactor - rateHistoryFactor + gallonsRequestedFactor + companyProfitFactor)
		// Update the profitMargin in the PricingModule
		_, _, err = client.From("Pricing Module").Update(map[string]interface{}{
			"profit_margin": margin,
		}, "", "").Eq("quote_id", quote_id).Execute()

		// Calculate the suggested price per gallon
		suggested_price_per_gallon := pricePerGallon + margin
		// Update the suggested_price in the FuelQuote
		_, _, err = client.From("Fuel Quote").Update(map[string]interface{}{
			"suggested_price": suggested_price_per_gallon,
		}, "", "").Eq("quote_id", quote_id).Execute()

		// Calculate the Total Amount Due
		totalAmountDue := float32(fuelQuote[0].GallonsRequested) * suggested_price_per_gallon
		// Update the TotalAmountDue in the FuelQuote
		_, _, err = client.From("Fuel Quote").Update(map[string]interface{}{
			"total_amount_due": totalAmountDue,
		}, "", "").Eq("quote_id", quote_id).Execute()

		crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.PricingModule]{
			Success: true,
			Data:    data,
		})
	}
}
