package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 4, -1, 1}
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j+1] < numbers[j] {
				temp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	sortedArray := numbers
	var result int
	for i := 0; i < n-1; i++ {
		if sortedArray[i+1] == sortedArray[i]+1 {
			continue
		} else if sortedArray[i]+1 > 0 {
			result = sortedArray[i] + 1
			break
		} else {
			continue
		}
	}
	if result == 0 {
		recurse(sortedArray[n-1] + 1)
	} else {
		fmt.Println(result)
	}
}

func recurse(count int) {
	if count > 0 {
		fmt.Println(count)
	} else {
		recurse(count + 1)
	}
}
