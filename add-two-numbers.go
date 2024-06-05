package main

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Num := ""
	l2Num := ""

	nextL1 := l1
	nextL2 := l2

	for nextL1 != nil {
		l1Num = strconv.Itoa(nextL1.Val) + l1Num
		nextL1 = nextL1.Next
	}

	for nextL2 != nil {
		l2Num = strconv.Itoa(nextL2.Val) + l2Num
		nextL2 = nextL2.Next
	}

	a, _ := new(big.Int).SetString(l1Num, 10)
	b, _ := new(big.Int).SetString(l2Num, 10)
	total := new(big.Int).Add(a, b).String()

	result := &ListNode{}
	current := result

	for i := len(total) - 1; i >= 0; i-- {
		current.Val, _ = strconv.Atoi(string(total[i]))

		if i > 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return result
}

func main() {}
