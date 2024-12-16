package main

import (
	"fmt"
	"sort"
)

// bit 全探索で解いたが、 N や K の範囲が大きいのでテストケースによってはエラー
func main() {
	var N, L int
	_, _ = fmt.Scan(&N, &L)

	var kireme int
	_, _ = fmt.Scan(&kireme)

	var pos = make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&pos[i])
	}

	var ans int
	// i の bit が立っている場合、i番目の切れ目で切り分ける
	for i := 0; i < (1<<N - 1); i++ {
		var kpos []int
		for j := 0; j < N-1; j++ {
			if i&(1<<j) > 0 {
				kpos = append(kpos, pos[j])
			}
		}

		if len(kpos) == kireme {
			m := minPartsLength(L, kpos)
			if ans < m {
				ans = m
			}
		}
	}

	fmt.Println(ans)
}

func minPartsLength(L int, kpos []int) int {
	var partsLens []int
	partsLenSum := 0

	for i := 0; i < len(kpos); i++ {
		l := kpos[i] - partsLenSum
		partsLens = append(partsLens, l)
		partsLenSum += l
	}

	partsLens = append(partsLens, L-partsLenSum)
	sort.Ints(partsLens)
	return partsLens[0]
}
