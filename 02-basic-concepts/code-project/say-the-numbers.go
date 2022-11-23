package main

import "fmt"

func main() {
	n := map[int]string{0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten"}
	var input [3]int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&input[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Println(n[input[i]])
	}

}
