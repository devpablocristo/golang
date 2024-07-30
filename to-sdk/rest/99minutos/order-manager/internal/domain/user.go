package domain

const (
	INTERNAL = iota + 1
	CUSTOMER
)

type User struct {
	UUID      string
	Username  string
	Email     string
	Password  string
	Role      int16
	CreatedAt int64
}
