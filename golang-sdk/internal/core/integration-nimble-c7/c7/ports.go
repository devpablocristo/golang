package cin7

type RedisPort interface {
	SaveShipment(Shipment) error
}
