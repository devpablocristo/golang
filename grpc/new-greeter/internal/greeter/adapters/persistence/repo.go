package repo

import portct "github.com/devpablocristo/qh/internal/greeter/ports"

type Repo struct {
	dao PortMongoGreetDAO
}

func NewRepo(dao PortMongoGreetDAO) portct.Repo {
	return &Repo{
		dao: dao,
	}
}
