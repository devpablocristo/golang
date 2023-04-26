package main

import "fmt"

type ErrPer struct {
	message string
	code    int
}

func NewErr(message string, code int) *ErrPer {
	return &ErrPer{
		message: message,
		code:    code,
	}
}

func (e *ErrPer) Error() string {
	return fmt.Sprintf("%s Error with code %d", e.message, e.code)
}

func checkErr(word string) (string, error) {
	if word == "error1" {
		return "", NewErr("Este es el error personalizado 1!", 1)
	} else if word == "error2" {
		return "", NewErr("Este es el error personalizado 2!", 2)
	}

	return "no hay error", nil
}

func main() {
	_, err := checkErr("error1")
	if err != nil {
		fmt.Println(err)
	}

	_, err = checkErr("error2")
	if err != nil {
		fmt.Println(err)
	}
}
