package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func nextIntSplit() []int {
	splitted := strings.Fields(nextLine())
	n := len(splitted)
	res := make([]int, n)
	for i, str := range splitted {
		res[i], _ = strconv.Atoi(str)
	}
	return res
}

func main() {
	nextInt()
	A := nextIntSplit()
	var alice, bob int
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i, a := range A {
		if i == 0 || i%2 == 0 {
			alice += a
		} else {
			bob += a
		}
	}

	fmt.Println(alice - bob)
}
