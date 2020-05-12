package main

import (
	"fmt"
	"os"
)

func main() {
	numbers := []int{10, 15, 3, 7}
	k := 17
	length := len(numbers)
	bruteForce(numbers, k, length)
	efficientSolution(numbers, k, length)
}

func bruteForce(numbers []int, k int, length int) {
	for i := 0; i < length; i++ {
		baseNum := numbers[i]
		for j := i + 1; j < length; j++ {
			if k == baseNum+numbers[j] {
				fmt.Println("true")
				os.Exit(0)
			}
		}
	}
	fmt.Println("false")
}

func efficientSolution(numbers []int, k int, length int) {
	dict := make(map[int]int, 0)
	for i := 0; i < length; i++ {
		if _, ok := dict[numbers[i]]; ok {
			fmt.Println("true")
			os.Exit(0)
		} else {
			dict[k-numbers[i]] = numbers[i]
			fmt.Println(dict)
		}
	}
	fmt.Println("false")
}
