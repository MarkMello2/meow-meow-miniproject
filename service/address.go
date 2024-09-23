package service

type AddressRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	Type      int    `json:"type"`
}

type AddressResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	Type      int    `json:"type"`
}

type AddressService interface {
	GetAddressByUserId(int) ([]AddressResponse, error)
	CreateAddress(AddressRequest, int, int) error
	DeleteAddressById(int) error
}
