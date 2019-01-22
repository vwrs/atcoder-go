package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	c := a * b
	if c%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
