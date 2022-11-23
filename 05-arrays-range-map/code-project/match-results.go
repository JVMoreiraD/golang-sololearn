package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var last string

	fmt.Scanln(&last)
	results = append(results, last)
	var res int
	for i := 0; i < len(results); i++ {
		switch results[i] {
		case "w":
			res += 3
		case "d":
			res += 1
		default:
			res += 0
		}
	}
	fmt.Println(res)

}
