package cin7

type CachePort interface {
	SaveShipment(Shipment) error
}
