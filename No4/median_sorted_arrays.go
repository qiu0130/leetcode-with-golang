package main



func binary_search_first_pos(nums []int, low, high, target int) int {

	var mid int
	l := low
	h := high

	for {
		mid = int((l + h) / 2)
		if target >= nums[mid] {
			return mid
		}
		l = mid
		if l == high-1 {
			return -1
		}
	}
}
func binary_search_last_pos(nums []int, low, high, target int) int {

	var mid int
	l := low
	h := high

	for {
		mid = int((l + h) / 2)
		if target <= nums[mid] {
			return mid
		}
		l = mid
		if l == high-1 {
			return -1
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	k := n + m
	mid := int(k/2)

	firstPos := binary_search_first_pos(nums1, 0, n, nums2[0])
	if firstPos == -1 {
		if k%2 == 0 {
			if mid <= n {
				if mid+1 <= n {
					ans := (nums1[mid] + nums1[mid+1]) / 2
				} else {
					ans := (nums1[mid] + nums2[0]) / 2
				}
			} else {
				pos := mid - n + 1
				ans := (nums2[pos] + nums2[pos+1]) / 2
			}
		} else {
			if mid <= n {
				ans := nums1[mid]
			} else {
				pos := n-mid+1
				ans := nums2[pos]
			}
		}
	} else {
		lastPos := binary_search_last_pos(nums1, 0, n, nums2[m-1])
		if lastPos == -1 {
			if k%2 == 0 {
				if mid <= n {

				} else {

				}
			} else {

			}
		} else {

		}
	}



}

/*
 1 2 3 4
 4 5
l=0, h=3 mid=1 nums[mid]=2
l=1, h=3, mid=2 nums[mid]=3
1 4 6
5 7
l=0 h=3 mid=1 nums[mid]=4
mid=1
l=0, h=3, mid=1, nums[mid]=4
l=1 h=3, mid=2 nums[mid]=8
mid=2


14578
 f-s

k=(n+m)/2
if k <
0 1 2 3 4 5
0 1 1 2 1 2


 */
func main() {

}
