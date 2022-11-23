package main

import "fmt"

func main() {
	var years int
	fmt.Scanln(&years)

	//your code goes here
	initialRabbits := 7

	for i := 0; i < years; i++ {
		initialRabbits = initialRabbits * 2
	}

	fmt.Println(initialRabbits)
}
