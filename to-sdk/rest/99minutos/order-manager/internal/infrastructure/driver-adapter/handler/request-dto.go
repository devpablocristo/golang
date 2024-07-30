package handler

import (
	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

type ShippingRequest struct {
	CustomerID  string    `json:"customer_id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	OrigAddress Address   `json:"orig_address"`
	DestAddress Address   `json:"dest_address"`
	Products    []Product `json:"products"`
}

type Address struct {
	Coords   Coords `json:"coords"`
	Zipcode  string `json:"zipcode"`
	Street   string `json:"address"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	ExtNum   string `json:"ext_num"`
	IntNum   string `json:"int_num"`
}

type Coords struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Product struct {
	Quantity int16   `json:"quantity"`
	Weight   float32 `json:"weight"`
	Size     int16   `json:"size"`
}

func (request *ShippingRequest) toOrderDomain() *domain.Order {
	var order domain.Order
	order.CustomerID = request.CustomerID
	order.OrigAddress = convertAddress(request.OrigAddress)
	order.DestAddress = convertAddress(request.DestAddress)
	order.Products = convertProducts(request.Products)

	return &order
}

func convertAddress(address Address) domain.Address {
	return domain.Address{
		Coords: domain.Coords{
			Latitude:  address.Coords.Latitude,
			Longitude: address.Coords.Longitude,
		},
		Zipcode:  address.Zipcode,
		Street:   address.Street,
		City:     address.City,
		Province: address.Province,
		Country:  address.Country,
		ExtNum:   address.ExtNum,
		IntNum:   address.IntNum,
	}
}

func convertProducts(reqProducts []Product) []domain.Product {
	products := make([]domain.Product, len(reqProducts))
	for i, p := range reqProducts {
		products[i] = domain.Product{
			Quantity: p.Quantity,
			Weight:   p.Weight,
			Size:     p.Size,
		}
	}
	return products
}
