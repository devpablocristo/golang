package portspkg


type MapDbClient interface {
	GetDb() map[string]interface{}
}
