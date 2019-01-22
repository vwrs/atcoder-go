package main

import (
	"fmt"
)

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

func reverseStr(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	arr := []int{1, 1, 2, 2, 2, 3, 4, 5}
	fmt.Println(uniqInt(arr))
}
