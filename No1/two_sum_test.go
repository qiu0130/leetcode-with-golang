package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	var tests = []struct {
		nums []int
		target int
		want []int
	}{
		{
			[]int {2, 7, 11, 15},
			9,
			 []int {0, 1},
		},
	}

	for _, test := range tests {
		if got := TwoSum(test.nums, test.target); got != nil {
			t.Errorf("got values [%q], wanted %v", got, test.want)
		}
	}
}