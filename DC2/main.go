package main

import "fmt"

func main() {
	numbers := []int{10, 3, 5, 6, 2}
	n := len(numbers)
	withDivision(numbers, n)
	withoutDivision(numbers, n)
}

func withDivision(numbers []int, n int) {
	product := 1
	for i := 0; i < n; i++ {
		product = product * numbers[i]
	}
	var result [5]int
	for i := 0; i < n; i++ {
		result[i] = product / numbers[i]
	}
	fmt.Println(result)
}

func withoutDivision(numbers []int, n int) {
	var left, right, product [5]int
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * numbers[i-1]
	}
	for j := n - 2; j >= 0; j-- {
		right[j] = right[j+1] * numbers[j+1]
	}
	for k := 0; k < n; k++ {
		product[k] = left[k] * right[k]
	}
	fmt.Println(product)
}
