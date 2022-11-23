package main

import "fmt"

func main() {
	team := map[string]float32{
		"P1": 1.98,
		"P2": 2.05,
		"P3": 1.89,
		"P4": 2.0,
		"P5": 2.11}
	var avg float32
	var count int
	for _, v := range team {
		count, avg = count+1, avg+v
	}
	fmt.Println(avg / float32(count))

}
