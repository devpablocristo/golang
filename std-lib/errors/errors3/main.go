package main

import (
	"errors"
	"fmt"
)

func addTwo(value int) (int, error) {
	if value < 0 {
		// business rules .. I can not work with negatives
		return value, errors.New(fmt.Sprintf("something went wrong value is %d:", value))
	} else {
		return value + 2, nil
	}
}

func main() {

	v, err := addTwo(-1)
	if err != nil {
		// do something with this error
		fmt.Print(err)
	} else {
		fmt.Print(v)
	}

}
