package handler

import "fmt"

func Chk(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintln(msg, err))
	}
}
