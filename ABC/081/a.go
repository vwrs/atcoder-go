package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	count := 0
	for _, char := range s {
		v, _ := strconv.Atoi(string(char))
		count += v
	}
	fmt.Println(count)
}
