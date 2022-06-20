package main

import "fmt"

func main() {
	// for {} // while true

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v)
	}
}
