package main

import (
	"fmt"
	"math/rand/v2"
	"math"
	"sort"
)

func binary_search(list []int, item int) int {
	var low int = 0
	var high int = len(list) - 1
	fmt.Printf("item: %d;\n", item)

	for low <= high {
		var half int = (low + high) / 2
		var guess int = list[half]

		fmt.Printf("guess = %d; half = %d\n",guess, half)
		if guess == item {
			return half
		} else if guess > item {
			high = half - 1
		} else {
			low = half + 1
		}
	}
	return 0
}

func main() {
	var i int = 0
	var size int = int(math.Pow(2, 20))
	var my_list []int
	// fmt.Printf("size = %d\n",size)
	for i < size {
		var random_number int = rand.IntN(1000000)
		// fmt.Printf("[%d] number = %d\n", i, random_number)
		my_list = append(my_list, random_number)
		i++
	}
	sort.Ints(my_list)
	// fmt.Printf("my_list = %d\n", my_list)
	var random_number int = rand.IntN(1000000)
	// fmt.Printf("number = %d\n", random_number)
	fmt.Println(binary_search(my_list, random_number))
}