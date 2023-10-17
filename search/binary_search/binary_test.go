package binarysearch

import (
	"fmt"
	"testing"
)

/*
必须采用顺序存储结构 2.必须按关键字大小有序排列
*/

func BinarySearch(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return mid
		} else if k < nums[mid] {
			right = mid - 1
		} else if k > nums[mid] {
			left = mid + 1
		}
	}
	return -1
}

func TestBS(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 3
	res := BinarySearch(nums, k)
	fmt.Println(res)
}

//todo
//插值查找
//基于二分查找算法，将查找点的选择改进为自适应选择，可以提高查找效率。它是二分查找的改进版。
//
