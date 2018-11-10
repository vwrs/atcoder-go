package main

import (
	"bufio"
	"fmt"
	"math/big"
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

var zero = big.NewInt(0)

func gcd(a, b *big.Int) *big.Int {
	for b.Cmp(zero) != 0 {
		t := b
		b = a.Mod(a, b)
		a = t
	}
	return a
}

func lcm(arr ...big.Int) *big.Int {
	if len(arr) < 2 {
		panic("error")
	}
	mul := big.NewInt(0)
	result := big.NewInt(0)
	mul.Mul(&arr[0], &arr[1])
	result.Div(mul, gcd(&arr[0], &arr[1]))

	for i := 2; i < len(arr); i++ {
		result = lcm(*result, arr[i])
	}

	return result
}

func main() {
	N := nextInt()
	if N == 1 {
		res := big.NewInt()
		res.SetString(nextLine(), 10)
		fmt.Println(res)
	} else {
		arr := make([]big.Int, N)
		i := 0
		for i < N {
			arr[i].SetString(nextLine(), 10)
			i++
		}
		fmt.Println(lcm(arr...))
	}
}
