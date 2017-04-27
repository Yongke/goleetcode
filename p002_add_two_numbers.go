package goleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersAux(l1, l2, 0)
}

func addTwoNumbersAux(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{carry, nil}
	}

	if l1 == nil && l2 != nil && carry == 0 {
		return l2
	}

	if l1 != nil && l2 == nil && carry == 0 {
		return l1
	}

	newVal := valOf(l1) + valOf(l2) + carry
	newCarry := 0
	if newVal >= 10 {
		newCarry = 1
		newVal -= 10
	}

	return &ListNode{newVal, addTwoNumbersAux(nextOf(l1), nextOf(l2), newCarry)}
}

func valOf(n *ListNode) int {
	if n == nil {
		return 0
	}
	return n.Val
}

func nextOf(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}
