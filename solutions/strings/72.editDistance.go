package strings

/* 72. Edit Distance
https://web.stanford.edu/class/cs124/lec/med.pdf

Defining dp map:
dp[i][j] is the distance from str1[1:i] to str2[1:j]
*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	// generate matrix in golang
	dp := [][]int{}
	// initialize the mapping
	for i := range word1 {
		dp[i][0] = i
	}
	for j := range word2 {
		dp[0][j] = j
	}

	// levenshtein algo
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			tempMin := min(dp[i-1][j]+1, dp[i][j-1]+1)
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(tempMin, dp[i-1][j-1])
			} else {
				dp[i][j] = min(tempMin, dp[i-1][j-1]+2)
			}
		}
	}
	return dp[n-1][m-1]
}
