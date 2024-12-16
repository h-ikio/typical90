package main

import "fmt"

func main() {
	var N, L int
	_, _ = fmt.Scan(&N, &L)

	var K int
	_, _ = fmt.Scan(&K)

	var pos = make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&pos[i])
	}

	// スコアを対象にして二分探索
	// 各スコアに対して、スコアを満たしつつ K + 1 個以上のパーツに切り分けられるかをチェックする
	// スコアを満たしつつ K + 1 個以上に切り分けられるということは、スコアを満たしつつちょうど K + 1 個にすることも可能なので、"以上" であることを確認すれば十分
	low, high := -1, L+1
	score := (low + high) / 2

	for {
		partsNum := 0
		prevK := 0
		for _, p := range pos {
			if p-prevK >= score {
				partsNum++
				prevK = p
			}
		}
		if L-prevK >= score {
			partsNum++
		}

		if partsNum >= K+1 {
			low = score
		} else {
			high = score
		}
		score = (low + high) / 2

		if high-low <= 1 {
			fmt.Println(low)
			break
		}
	}
}
