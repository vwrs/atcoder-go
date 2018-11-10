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
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

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

func nextStrSplit() []string {
	splitted := strings.Fields(nextLine())
	n := len(splitted)
	res := make([]string, n)
	for i, str := range splitted {
		res[i] = str
	}
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
	defer writer.Flush()
	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("receive %d and %d.\n", a, b)
	// input := nextLine()
	// fmt.Println(nextIntSplit())
}
