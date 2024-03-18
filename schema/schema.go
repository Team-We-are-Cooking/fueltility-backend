package schema

type User struct {
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	AddressTwo string `json:"address_two"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zip_code"`
}

type FuelQuote struct {
	QuoteId          int8    `json:"quote_id"`
	Username         string  `json:"username"`
	Interstate       bool    `json:"interstate"`
	DeliveryAddress  string  `json:"delivery_address"`
	DeliveryDate     string  `json:"delivery_date"`
	GallonsRequested int8    `json:"gallons_requested"`
	SuggestedPrice   float32 `json:"suggested_price"`
	TotalAmountDue   float32 `json:"total_amount_due"`
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type PricingModule struct {
	QuoteId 	int8    `json:"quote_id"`
	ProfitMargin float32 `json:"profit_margin"`
	CalculatedTotalCost float32 `json:"calculated_total_cost"`
}