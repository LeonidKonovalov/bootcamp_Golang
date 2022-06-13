package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
		fallthrough
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}
	// 	switch 2 {
	// case 1:
	//     fmt.Println("1")
	//     fallthrough
	// case 2:
	//     fmt.Println("2")
	//     fallthrough
	// case 3:
	//     fmt.Println("3")
	// }
	// Loop:
	// 	for _, ch := range "a b\nc" {
	// 		switch ch {
	// 		case ' ':
	// 			break
	// 		case '\n':
	// 			break Loop
	// 		default:
	// 			fmt.Printf("%c\n", ch)
	// 		}
	// 	}
}
