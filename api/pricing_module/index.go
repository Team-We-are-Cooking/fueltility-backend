package pricing_module

import (
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
		quote_id := r.URL.Query().Get("quote_id")
		if quote_id == "" {
			crw.SendJSONResponse(http.StatusBadRequest, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: "Missing quote id."},
			})
			return
		}

		var data schema.PricingModule
		if _, err := client.From("Pricing Module").Select("*", "exact", false).Eq("quote_id", quote_id).ExecuteTo(&data); err != nil {
			crw.SendJSONResponse(http.StatusInternalServerError, fueltilityhttp.ErrorResponse{
				Success: false,
				Error:   &fueltilityhttp.ErrorDetails{Message: err.Error()},
			})
			return
		}

		// crw.SendJSONResponse(http.StatusOK, fueltilityhttp.Response[schema.PricingModule]{
		// 	Success: true,
		// 	Data:    data,
		// })
}
}