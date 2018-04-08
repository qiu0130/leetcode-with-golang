
package two_sum

func TwoSum(nums []int, target int) []int {

	for i, a := range nums {
		for j, b := range nums {
				if i != j {
					if a+b == target {
						return []int {i, j}
					}
				}
		}
	}
	return []int{}
}
