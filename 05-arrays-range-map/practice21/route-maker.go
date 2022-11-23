package main

import "fmt"

//create the route() function
func route(cities ...string) {
	var route string

	for _, v := range cities {
		route += v
		route += "->"
	}

	fmt.Println(route)
}
func main() {
	var n int
	fmt.Scanln(&n)

	var cities []string
	//take n strings as input and append them to the slice
	for i := 0; i < n; i++ {
		var input string
		fmt.Scanln(&input)

		cities = append(cities, input)
	}
	//
	route(cities...)
}
