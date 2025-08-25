package arrays

import "fmt"

func removeDuplicates(nums []int) int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[x] {
			x += 1
			nums[x] = nums[i]
		}
	}
	return x + 1
}

func RemoveDuplicatesTest() {
	numberOfUniqueElements := removeDuplicates([]int{1, 1, 2})
	fmt.Println(numberOfUniqueElements) // Output: 2
	numberOfUniqueElements = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(numberOfUniqueElements) // Output: 5
}
