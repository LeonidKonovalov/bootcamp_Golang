package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
		fallthrough // next case
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}

Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ':
			break
		case '\n':
			break Loop // переход по метке
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}
