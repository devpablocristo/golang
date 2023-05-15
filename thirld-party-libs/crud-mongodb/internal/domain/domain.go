package domain

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
	IsExact     bool      `bson:"is_location_exact"`
}

type Address struct {
	Street         string   `bson:"street"`
	Suburb         string   `bson:"suburb"`
	GovernmentArea string   `bson:"government_area"`
	Market         string   `bson:"market"`
	Country        string   `bson:"country"`
	CountryCode    string   `bson:"country_code"`
	Location       Location `bson:"location"`
}

type Listing struct {
	ID      string  `bson:"_id"`
	Name    string  `bson:"name"`
	Address Address `bson:"address"`
	Summary string  `bson:"summary"`
}
