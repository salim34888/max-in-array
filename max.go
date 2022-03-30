package main

import "fmt"

func main() {
	var array []int
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		array = append(array, number)
	}

	max := array[0]
	for _, element := range array {
		if element > max {
			max = element
		}
	}

	fmt.Println(max)
}
