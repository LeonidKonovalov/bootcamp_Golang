package main

import (
	"errors"
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	err := errors.New("hello")
	if err != nil {
		panic(err)
	}

}
