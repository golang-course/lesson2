package main

import "fmt"

func main() {
	var a int
	a = 3
	switch {
	case a > 1:
		fmt.Println("а > 1")
		fallthrough
	case a == 100:
		fmt.Println("3")
	case a == 4 || a == 5:
		fmt.Println("5 or 4")
	default:
		fmt.Println("Default case")
	}
}
