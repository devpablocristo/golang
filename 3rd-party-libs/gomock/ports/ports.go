package ports

import service "github.com/JoakoMeLi/go-concepts/3rd-party-libs/gomock"

type Accessor interface {
	Save(n int, p service.Person)
	Retrieve(n int) service.Person
}
