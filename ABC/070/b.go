package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func main() {
	defer writer.Flush()

	var a, b, c, d int
	scanf("%d %d %d %d\n", &a, &b, &c, &d)
	common := int(math.Min(float64(b), float64(d)) - math.Max(float64(a), float64(c)))
	if common < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(common)
	}
}
