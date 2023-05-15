package domain

type Listing struct {
	Name    string `bson:"name"`
	Address string `bson:"address"`
	City    string `bson:"city"`
}
