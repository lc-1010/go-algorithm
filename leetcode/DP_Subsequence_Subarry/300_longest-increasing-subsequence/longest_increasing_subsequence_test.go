package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/

func Test_lengthOfLIS(t *testing.T) {
	l := lengthOfLIS([]int{1, 31, 4, 5, 10, 13})
	assert.Equal(t, 5, l)
}

// i【 1, 31, 4, 5, 10, 13 】
//
//	j【 1, 31, 4, 5, 10, 13 】
//
// 从头开始对比
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		// 初始化为1
		dp[i] = 1 //每一个都是自己的一个子序列
		//同时不需要连续，所有只需要连续的比较就好了
	}
	// 获取一个基准位作为最大的比较
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				// 和之前的对比一下取最大值
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)

	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}
	fmt.Printf("%v", dp)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	fmt.Printf("\n%v\n", dp)
	return dp[n][m]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	fmt.Printf("%v", dp)
	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {

			if word1[m-1] == word2[n-1] {
				dp[i][j] = dp[i-1][j-1]

			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)

			}

		}

	}
	fmt.Printf("\n%v\n", dp)

	return dp[m][n]
}
func TestDel(t *testing.T) {
	print(minDistance1("sea", "eat"))
	print("\n")
	print(minDistance("sea", "eat"))

}
