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

type coord struct {
	y, x, step int
}

func main() {
	var h, w int
	fmt.Scanf("%d %d\n", &h, &w)

	field := make([][]string, h)
	for i := range field {
		field[i] = nextRuneSplit()
	}

	steps := make([][]int, h)
	for i := range steps {
		steps[i] = make([]int, w)
	}

	// count black tiles
	nBlack := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if field[y][x] == "#" {
				nBlack++
			}
			steps[y][x] = -1
		}
	}

	// BFS
	queue := []coord{}
	queue = append(queue, coord{0, 0, 0})
	steps[0][0] = 0 // init

	moves := [][]int{
		[]int{-1, 0}, // left
		[]int{1, 0},  // right
		[]int{0, -1}, // up
		[]int{0, 1},  // down
	}

	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:] // deque
		for _, s := range moves {
			dx, dy := s[0], s[1]
			x := now.x + dx
			y := now.y + dy

			// field && unvisited
			if y >= 0 && y < h && x >= 0 && x < w && steps[y][x] == -1 && field[y][x] == "." {
				queue = append(queue, coord{y, x, now.step + 1})
				steps[y][x] = now.step + 1
			}
		}
	}

	// result
	if steps[h-1][w-1] == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(h*w - nBlack - (steps[h-1][w-1] + 1))
	}
}
