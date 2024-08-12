package core

import (
	"context"
)

type GreaterUseCasesPort interface {
	Hello(context.Context) (string, error)
}

type greeterUseCases struct{}

func NewGreaterUseCases() GreaterUseCasesPort {
	return &greeterUseCases{}
}

func (s *greeterUseCases) Hello(ctx context.Context) (string, error) {
	return "hello", nil
}
