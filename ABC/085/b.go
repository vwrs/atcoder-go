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

func uniqInt(arr []int) []int {
	uniq := []int{}
	m := make(map[int]bool)
	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}

func main() {
	var N int
	fmt.Scan(&N)
	D := make([]int, N)
	for i := range D {
		D[i] = nextInt()
	}
	uniq := uniqInt(D)
	fmt.Println(len(uniq))
}
