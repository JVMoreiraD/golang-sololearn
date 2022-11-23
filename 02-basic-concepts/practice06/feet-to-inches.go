package main

import "fmt"

func main() {
	//your code goes here
	var feet int
	var inches int

	//scan feet
	fmt.Scanln(&feet)

	//new inches
	inches = feet * 12

	fmt.Println(inches)

}
