/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rootNode *ListNode
	var currentNode *ListNode

	isInitial := true
	carryOver := 0

	for {
		co, res := addNode(carryOver, l1, l2)
		resNode := ListNode{
			Val:  res,
			Next: nil,
		}

		carryOver = co

		if isInitial {
			rootNode = &resNode
			currentNode = rootNode
			isInitial = false
		} else {
			currentNode.Next = &resNode
			currentNode = currentNode.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if carryOver == 1 {
				currentNode.Next = &ListNode{Val: carryOver}
			}

			break
		}
	}

	return rootNode
}

func addNode(co int, l1 *ListNode, l2 *ListNode) (int, int) {
	res := co

	if l1 != nil && l2 != nil {
		res += l1.Val
		res += l2.Val
	} else if l1 == nil {
		res += l2.Val
	} else if l2 == nil {
		res += l1.Val
	}

	if res > 9 {
		return 1, res - 10
	} else {
		return 0, res
	}
}