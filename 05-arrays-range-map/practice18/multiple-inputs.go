package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	//your code goes here
	array := make([]int, 0)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanln(&x)
		array = append(array, x)
	}

	fmt.Println(array)
}
