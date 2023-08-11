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
