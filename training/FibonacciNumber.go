package training

import (
	"fmt"
)

func fibonacciNumber(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	numbers := make([]int, n)
	numbers[0], numbers[1] = 1, 1

	for i := 2; i < n; i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}

	return numbers
}

func main() {
	var num int

	fmt.Print("number: ")
	fmt.Scanln(&num)
	fmt.Println(fibonacciNumber(num))
}
