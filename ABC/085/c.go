package main

import (
	"bufio"
	"fmt"
	"os"
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
	res := nextIntSplit()
	N := res[0]
	Y := res[1]
	var z int
	xans, yans, zans := -1, -1, -1
	for x := 0; x <= N; x++ {
		for y := 0; x+y <= N; y++ {
			z = N - x - y
			if 10000*x+5000*y+1000*z == Y {
				xans, yans, zans = x, y, z
				goto ANS
			}
		}
	}

ANS:
	fmt.Println(xans, yans, zans)
}
