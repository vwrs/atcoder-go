package main

import (
	"fmt"
	"strconv"
)

func sumDigit(s string) int {
	sum := 0
	for _, char := range s {
		v, _ := strconv.Atoi(string(char))
		sum += v
	}
	return sum
}

func main() {
	var N, A, B, tmpSum, sum int
	fmt.Scan(&N, &A, &B)
	for i := 1; i <= N; i++ {
		tmpSum = sumDigit(strconv.Itoa(i))
		if tmpSum >= A && tmpSum <= B {
			sum += i
		}
	}
	fmt.Println(sum)
}
