package service

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func funcTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	acc := NewMockAccessor(ctrl)

	p := Person{
		First: "James",
	}

	//set the expectation
	//extpect save this parameters min 1 time, max 1 time
	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)
}
