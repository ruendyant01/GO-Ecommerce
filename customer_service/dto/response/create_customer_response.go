package response

type CreateCustomerResponse struct {
	CustomerId string         `json:"customer_id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Address    CompanyAddress `json:"address"`
}

type CompanyAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
}
