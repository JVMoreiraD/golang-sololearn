package main

import "fmt"

func mars_age(x int) int {
	var days int = x * 365
	var newAge int = (days / 687)

	return (newAge)

}
func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
