package main

import (
	"errors"
	"fmt" // форматированный ввод и вывод
)

func main() { // точка входа
	var i int // default value zero
	j := 1    // short assing, var j int = 1
	a := [5]int{1, 2, 3, 4, 5}
	// массив, без добавления
	fmt.Println(a)
	s := []int{1, 2, 3, 4, 5}
	// slice, динамический массив
	fmt.Println(s)
	err := errors.New("hello")
	if err != nil {
		panic(err)
	}

}
