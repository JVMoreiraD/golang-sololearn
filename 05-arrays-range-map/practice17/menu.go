package main

import "fmt"

func main() {
	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	//your code goes here
	var input int

	fmt.Scanln(&input)

	if input > 6 || input < 0 {
		fmt.Println("Invalid choice")
	} else {
		fmt.Println(menu[input])
	}
}
