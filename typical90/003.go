package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var ab = make([][]int, N)

	for i := 0; i < N-1; i++ {
		a := nextInt(sc) - 1
		b := nextInt(sc) - 1
		ab[a] = append(ab[a], b)
		ab[b] = append(ab[b], a)
	}

	// 頂点0 から最も遠い頂点と距離
	dist0 := dist(ab, 0, N)
	ri, rv := -1, -1
	for idx, v := range dist0 {
		if rv < v {
			ri = idx
			rv = v
		}
	}

	// 頂点i から最も遠い頂点と距離
	distI := dist(ab, ri, N)
	ret := 0
	for i := 0; i < N; i++ {
		if ret < distI[i] {
			ret = distI[i]
		}
	}

	fmt.Println(ret + 1)
}

// i から各頂点への距離を表す配列を返す
func dist(ab [][]int, i int, N int) []int {
	var d = make([]int, N)
	for idx := 0; idx < N; idx++ {
		d[idx] = -1
	}
	d[i] = 0

	s := make([]int, 0, N)
	s = append(s, i)

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		for _, t := range ab[v] {
			if d[t] == -1 {
				s = append(s, t)
				d[t] = d[v] + 1
			}
		}
	}

	return d
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
