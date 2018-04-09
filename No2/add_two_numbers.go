package main
import (
	//"fmt"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func Reverse(s []int) []int {
	r := make([]int, len(s))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		r[j] = s[i]
	}
	return r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {


	firstNumbers := []int {}
	secondNumbers := []int {}

	tmp := l1
	for {
		firstNumbers = append(firstNumbers, tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	tmp = l2
	for {
		secondNumbers = append(secondNumbers, tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}
	length := len(firstNumbers) - len(secondNumbers)
	if length > 0 {
		for i := 0; i < length; i = i+1 {
			secondNumbers = append(secondNumbers, 0)

		}
	} else {
		for i := 0; i < -length; i = i+1 {
			firstNumbers = append(firstNumbers, 0)

		}
	}

	sumNumbers := make([]int, len(firstNumbers)+1)
	sumNumbers[0] = 0
	firstNumbers = Reverse(firstNumbers)
	secondNumbers = Reverse(secondNumbers)

	//fmt.Println(firstNumbers)
	//fmt.Println(secondNumbers)

	for i := 1; i < len(firstNumbers)+1; i = i+1 {
		sumNumbers[i] = firstNumbers[i-1] + secondNumbers[i-1]
	}
	for i := len(firstNumbers); i > 0; i = i-1 {
		flag := sumNumbers[i]
		if flag >= 10 {
			sumNumbers[i-1] = sumNumbers[i-1] + 1
		}
		sumNumbers[i] = flag%10
	}
	if sumNumbers[0] == 0 {
		sumNumbers = sumNumbers[1:]
	}
	sumNumbers = Reverse(sumNumbers)
	//fmt.Println(sumNumbers)

	return initListNode(sumNumbers)
}

func initListNode(nums []int) *ListNode {

	if len(nums) == 0 {
		return nil
	}
	l := new(ListNode)
	l.Val = nums[0]
	l.Next = nil

	tmp := l
	for _, v := range nums[1:] {
		ll := new(ListNode)
		ll.Val = v
		ll.Next = nil

		tmp.Next = ll
		tmp = ll
	}
	//PrintListNode(l)
	return l
}

func PrintListNode(pre *ListNode) {
	tmp := pre
	for {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			fmt.Println()
			break
		} else {
			fmt.Print("->")
		}
	}
}

func main() {

	var tests = []struct {
		n1 []int
		n2 []int
	}{
		{
			[]int {2, 4, 3},
			[]int {5, 6, 4},
		},
		{
			[]int {5},
			[]int {5},
		},
		{
			[]int {5, 4},
			[]int {4, 6},
		},
		{
			[]int {0},
			[]int {0},
		},
		{
			[]int {9, 8},
			[]int {1},
		},
	}
	for _, test := range tests {
		l1 := initListNode(test.n1)
		l2 := initListNode(test.n2)
		addTwoNumbers(l1, l2)
	}

}
