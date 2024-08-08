package patients

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/person"
)

type Patient struct {
	ID           string
	PersonalInfo person.Person
	CreatedAt    time.Time
}
