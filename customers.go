package gozoko

type Customer struct {
	Id string `json:"id"`
}

type CustomerService interface {
	List() ([]Customer, error)
	Get() (Customer, error)
	DeleteMessages(customerId string) error
}

type CustomerServiceImpl struct {
	client *Client
}
