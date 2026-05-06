/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{
		Val: 0,
		Next: nil,
	}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0{
		sum := valOrZero(l1) + valOrZero(l2) + carry
		carry = sum / 10
		val := sum % 10

		current.Next = &ListNode{
			Val: val,
			Next: nil,
		}

		current = current.Next
		l1 = nextOrNil(l1)
		l2 = nextOrNil(l2)
	}

	fmt.Printf("dummy %v\n", dummy)

	return dummy.Next
}

func valOrZero(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return 0
}


func nextOrNil(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return nil
	}

	return node.Next
}