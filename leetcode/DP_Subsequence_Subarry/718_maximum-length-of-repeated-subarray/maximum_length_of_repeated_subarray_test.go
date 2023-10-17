package maximumlengthofrepeatedsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

示例 1：

输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
*/
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
	}
	res := 0
	for i := 1; i <= m; i++ { //两个数组逐个对比
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 相等则+1
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}

	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestFindLength(t *testing.T) {
	n1 := []int{1, 2, 3, 2, 1}
	n2 := []int{3, 2, 1, 4, 7}
	res := findLength(n1, n2)
	assert.Equal(t, 3, res)
}
