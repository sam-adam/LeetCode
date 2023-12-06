func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	shouldSwap := true

	var newRoot *ListNode

	n1 := head
	n2 := head.Next

	for shouldSwap {
		newN1 := swapNode(n1, n2)

		if newRoot == nil {
			newRoot = newN1
		} else {
			appendNode(newRoot, newN1)
		}

		if n2.Next != nil && n2.Next.Next != nil {
			n1 = n2.Next
			n2 = n2.Next.Next
			shouldSwap = true
		} else if n2.Next != nil {
			appendNode(newRoot, n2.Next)
            shouldSwap = false
		} else {
			shouldSwap = false
		}
	}

	return newRoot
}

func appendNode(root *ListNode, newNode *ListNode) {
	currNode := root

	for true {
		if currNode.Next == nil {
			currNode.Next = newNode

			return
		} else {
			currNode = currNode.Next
		}
	}
}

func swapNode(n1 *ListNode, n2 *ListNode) (*ListNode) {
	return &ListNode{
		Val: n2.Val,
		Next: &ListNode{
			Val: n1.Val,
			Next: nil,
		},
	}
}
