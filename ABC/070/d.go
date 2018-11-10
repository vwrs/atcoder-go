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
	i, _ := strconv.Atoi(nextLine())
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

type edge struct {
	to   int
	cost uint
}

func dfs(graph [][]edge, costs []uint, node int, prev int, totalCost uint) {
	paths := graph[node]
	for _, p := range paths {
		if p.to != prev {
			cost := totalCost + p.cost
			costs[p.to] = cost
			dfs(graph, costs, p.to, node, cost)
		}
	}
}

func main() {
	n := nextInt()

	graph := make([][]edge, n+1)
	for i := 0; i < n-1; i++ {
		arr := nextIntSplit()
		a, b, cost := arr[0], arr[1], uint(arr[2])
		graph[a] = append(graph[a], edge{b, cost})
		graph[b] = append(graph[b], edge{a, cost})
	}

	qk := nextIntSplit()
	q, k := qk[0], qk[1]

	costs := make([]uint, n+1)
	dfs(graph, costs, k, -1, 0)

	for i := 0; i < q; i++ {
		q := nextIntSplit()
		fmt.Println(costs[q[0]] + costs[q[1]])
	}
}
