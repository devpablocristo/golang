package cin7

type CacheRepository interface {
	SaveShipment(Shipment) error
}
