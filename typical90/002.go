package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	// bit が立っていない場合は (、立っている場合は ) とする
	for i := 0; i < (1 << N); i++ {
		var S string

		// 辞書順にしたいので最上位 bit から見る
		for j := N - 1; 0 <= j; j-- {
			if i&(1<<j) == 0 {
				S += "("
			} else {
				S += ")"
			}
		}

		if isValid(S) {
			fmt.Println(S)
		}
	}
}

func isValid(S string) bool {
	var score int
	for _, r := range S {
		s := string(r)
		switch s {
		case "(":
			score += 1
		case ")":
			score -= 1
		}

		if score < 0 {
			return false
		}
	}

	return score == 0
}
