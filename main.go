package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите введите возраст: ")
	fmt.Scanln(&age)
	fmt.Printf("Hello, %v, %v \n", name, age)
	/*
		if age >= 18 {
			fmt.Println("Вы совершеннолетний")
		} else if age <= 0 {
			fmt.Println("Вы точно обманываете)")
		} 
	*/

}
