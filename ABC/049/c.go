package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func reverseStr(s string) string {
	strslice := strings.Split(s, "")
	fmt.Println(strslice)
	sort.Sort(sort.Reverse(sort.StringSlice(strslice)))
	return strings.Join(strslice, "")
}

func main() {
	input := nextLine()
	fmt.Println(strings.Index(input, "dream"))
	fmt.Println(reverseStr(input))
}
