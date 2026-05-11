/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	// base 
	if (head == nil || head.Next == nil) {
		return head
	}

	new_head := reverseList(head.Next)
	
	head.Next.Next = head
	head.Next = nil

	return new_head
}
