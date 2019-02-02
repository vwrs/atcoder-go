package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func nextBigInt() *big.Int {
	res := big.NewInt(0)
	res.SetString(nextLine(), 10)
	return res
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

func nextStrSplit() []string {
	splitted := strings.Fields(nextLine())
	n := len(splitted)
	res := make([]string, n)
	for i, str := range splitted {
		res[i] = str
	}
	return res
}

func nextRuneSplit() []string {
	return strings.Split(nextLine(), "")
}

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	fmt.Printf("receive %d and %d.\n", a, b)
}
