package ginkgofw

type Person struct {
	Age int
}

func (p *Person) IsChild() bool {
	if p.Age < 18 {
		return true
	}

	return false
}
