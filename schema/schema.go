package schema

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
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
	ProfitMargin float32 `json:"profit_margin"`
	QuoteId 	int8    `json:"quote_id"`
}