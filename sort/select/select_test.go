package selecttest

import (
	"fmt"
	"testing"
)

func selectSort(nums []int) {
	n := len(nums)
	//初始化
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func TestSeelctSort(t *testing.T) {
	nums := []int{3, 4, 5, 2, 23, 76, 34}
	selectSort(nums)
	fmt.Println(nums)
	a := 10
	ab(a)
	fmt.Println(a)
}
func ab(a int) {
	a = a + 1
}
