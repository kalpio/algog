package main

import (
	"fmt"
)

func main() {
	arr := []int{30, 5, 20, 17, 34, 51, 87, 49, 7, 57, 73, 47, 100, 59, 60, 97, 62, 87, 5, 68, 70, 33, 89,
		96, 14, 47, 42, 83, 55, 5, 77, 70, 86, 79, 52, 95, 26, 64, 25, 35, 25, 80, 9, 21, 42, 45,
		88, 23, 30, 56, 13, 82, 1, 35, 99, 60, 85, 29, 16, 19, 25, 15, 18, 98, 48, 8, 67, 2, 7, 75,
		88, 82, 97, 27, 14, 1, 25, 17, 61, 99, 18, 82, 1, 35, 96, 78, 9, 46, 41, 23, 18, 37, 25,
		74, 24, 39, 41, 94, 97, 16}

	result := InsertionSort(arr)

	fmt.Println(result)
}

func InsertionSort(data []int) []int {
	j := 1
	for j, _ = range data {
		key := data[j]
		i := j - 1
		for i >= 0 && data[i] > key {
			data[i+1] = data[i]
			i = i - 1
		}
		data[i+1] = key
	}

	return data
}
