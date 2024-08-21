package core

import (
	"context"
)

type GreaterUseCases interface {
	Hello(context.Context) (string, error)
}

type greeterUseCases struct{}

func NewGreaterUseCases() GreaterUseCases {
	return &greeterUseCases{}
}

func (s *greeterUseCases) Hello(ctx context.Context) (string, error) {
	return "hello", nil
}
