package domain

const (
	SM int16 = iota + 1
	MD
	LG
)

const CREATED string = "created"
const COLLECTED string = "collected"
const IN_STATION string = "in_station"
const ON_ROUTE string = "on_route"
const DELIVERED string = "delivered"
const CANCELED string = "canceled"

type Address struct {
	Coords   Coords
	Zipcode  string
	Street   string
	City     string
	Province string
	Country  string
	ExtNum   string
	IntNum   string
}

type Product struct {
	Quantity int16
	Weight   float32
	Size     int16
}

type Order struct {
	UUID        string
	CustomerID  string
	OrigAddress Address
	DestAddress Address
	Products    []Product
	PackageType int16
	TotalWeight float32
	Status      string
	CreatedAt   int64
	UpdatedAt   int64
}

type Coords struct {
	Latitude  float64
	Longitude float64
}
