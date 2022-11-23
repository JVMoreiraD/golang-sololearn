package main

import "fmt"

func main() {
	//your code goes here
	var temperature float32
	const temperatureLimit float32 = 99.5

	// fmt.Println("Type a temperature: ")
	fmt.Scan(&temperature)

	if temperature > temperatureLimit {
		fmt.Println("Fever")
	} else {
		fmt.Println("Allowed")
	}
}
