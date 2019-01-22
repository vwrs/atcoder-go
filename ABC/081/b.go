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

func main() {
	var N, count int
	fmt.Scan(&N)
	A := nextIntSplit()
	for {
		for i := 0; i < N; i++ {
			if A[i]%2 != 0 {
				goto ANS
			}
			A[i] /= 2
		}
		count++
	}

ANS:
	fmt.Println(count)
}
