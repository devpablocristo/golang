package math_test

import (
	"testing"

	mock_math "github.com/braddle/concurrency/math/math_mock"
	"github.com/golang/mock/gomock"

	"github.com/braddle/concurrency/math"
)

func TestMultiplyChannel(t *testing.T) {
	in := make(chan math.Table, 3)
	in <- math.Table{Base: 5}
	in <- math.Table{Base: 10}
	in <- math.Table{Base: 20}

	out := math.MultiplyTables(in, 3)

	a := <-out
	b := <-out
	c := <-out

	if a.Base != 5 || a.Result != 15 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}

	if b.Base != 10 || b.Result != 30 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}

	if c.Base != 20 || c.Result != 60 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}
}

func TestMultiplyChannelRandomNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	in := make(chan math.Table, 3)
	in <- math.Table{Base: 5}
	in <- math.Table{Base: 10}
	in <- math.Table{Base: 20}

	random := mock_math.NewMockRandom(ctrl)
	random.EXPECT().GenerateInt().Return(3).Times(3)

	out := math.MultiplyTablesRandom(in, random)

	a := <-out
	b := <-out
	c := <-out

	if a.Base != 5 || a.Result != 15 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}

	if b.Base != 10 || b.Result != 30 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}

	if c.Base != 20 || c.Result != 60 {
		t.Logf("Base:\t%d", a.Base)
		t.Logf("Result:\t%d", a.Result)
		t.Error("Multiplication failed")
	}
}
