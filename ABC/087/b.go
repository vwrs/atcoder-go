package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	i, e := strconv.Atoi(nextLine())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	var A, B, C, X, sum, count int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &A, &B, &C, &X)
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				sum = 500*a + 100*b + 50*c
				if sum == X {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
