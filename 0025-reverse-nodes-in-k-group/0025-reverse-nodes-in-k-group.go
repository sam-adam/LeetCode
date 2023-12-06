func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k < 2 {
		return head
	}

	var prevLastNode *ListNode
	var headNode *ListNode

	currNode := head

	for {
		var nodes []*ListNode

		for i := 0; i < k; i++ {
			if currNode != nil {
				nodes = append(nodes, currNode)
				currNode = currNode.Next
			}
		}

		if len(nodes) < k {
			if prevLastNode != nil {
				prevLastNode.Next = nodes[0]
			} else if headNode == nil {
				headNode = head
			}

			break
		}

		for i := k - 1; i >= 0; i-- {
			if i == k - 1 {
				if prevLastNode != nil {
					prevLastNode.Next = nodes[i]
				}
			} else {
				nodes[i].Next = nil
				nodes[i+1].Next = nodes[i]

				if i == 0 {
					prevLastNode = nodes[i]
				}
			}
		}

		if headNode == nil {
			headNode = nodes[k - 1]
		}

		if currNode == nil {
			break
		}
	}

	return headNode
}
