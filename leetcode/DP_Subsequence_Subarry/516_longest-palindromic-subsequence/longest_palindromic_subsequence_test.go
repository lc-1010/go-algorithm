package longestpalindromicsubsequence

/*
516. 最长回文子序列

给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。
*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1 //初始化 一个1时都为1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//dp[i][j] 是s（i，j） 最大回文串
	// i---> xxx<--j
	// i+1 ,j-1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ { // 每一个自己都是一个一个回文
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}

		}
	}
	return dp[0][n-1]
}
