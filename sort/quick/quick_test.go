package quick

import (
	"fmt"
	"testing"
)

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	mid := 0
	left, right := 0, n-1
	for left <= right {
		if nums[left] > nums[mid] && nums[right] < nums[mid] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left] <= nums[mid] {
			left++
		}
		if nums[right] >= nums[mid] {
			right--
		}
	}
	nums[mid], nums[right] = nums[right], nums[mid]
	quickSort(nums[:right])
	quickSort(nums[right+1:])
}
func TestQuick(t *testing.T) {
	arr := []int{6, 2, 8, 1, 4, 9, 3, 7, 5}
	quickSort(arr)
	fmt.Println(arr)
}

func Pal(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	fmt.Printf("%v", dp)
	res := 0

	for i := len(s) - 1; i >= 0; i-- {

		for j := i; j < len(s); j++ {
			fmt.Println("i,j", i, j)
			// if s[i] == s[j] {
			// 	if j-i <= 1 {
			// 		dp[i][j] = true
			// 		res++
			// 	} else if dp[i+1][j-1] {
			// 		dp[i][j] = true
			// 		res++
			// 	}
			// }
		}
	}
	fmt.Printf("\n%v", dp)
	return res
}

func TestPal(t *testing.T) {
	s := "aaab"
	Pal(s)
}

/*
	[false false false false]
	[false false false false]
	[false false false false]
	[false false false false]

-----------------------

	[true true true false]
	[false true true false]
	[false false true false]
	[false false false true]
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PalMax(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1

	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][len(s)-1]
}

func TestPalMax(t *testing.T) {
	s := "cbbd"
	print(PalMax(s))
}
