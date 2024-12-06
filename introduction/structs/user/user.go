package user

type User struct {
	Name        string
	LastName    string
	Age         int
	UserAddress Address
}

type Address struct {
	Street string // accesible
	number int    // not accesible
}

// func Foo()
// func bar()
