package usr

type User struct {
	Username string
	Password string
}

type InMemDB map[string]*User
