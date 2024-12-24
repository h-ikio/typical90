package main

import (
	"fmt"
	"sort"
)

// DP を使っているがこの解き方だと必ずしも最長にならないケースがある
// 正確には現地点から一番遠い頂点を基準に、そこから最も遠い頂点を探して木の直径を求める必要がある
func main() {
	var N int
	_, _ = fmt.Scan(&N)

	var dp = make([]int, N+1)

	for i := 0; i < N-1; i++ {
		var a, b int
		_, _ = fmt.Scan(&a, &b)

		dp[b] = max(dp[a]+1, dp[b])
	}

	for i, _ := range dp {
		fmt.Println(i, dp[i])
	}
	sort.Ints(dp)
	fmt.Println(dp[len(dp)-1]+1, dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
