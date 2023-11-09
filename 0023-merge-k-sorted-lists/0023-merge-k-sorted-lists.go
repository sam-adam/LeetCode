func mergeKLists(lists []*ListNode) *ListNode {
	var resultNode *ListNode

	if lists == nil || len(lists) == 0 {
		return nil
	}

	for len(lists) > 0 {
		leastIdx := -1
		leastValue := math.MaxInt32

		for idx, l := range lists {
            if l == nil {
                continue
            }

			if l.Val < leastValue {
				leastValue = l.Val
				leastIdx = idx
			}
		}
        
        if leastIdx == -1 {
            return resultNode
        }

		if resultNode == nil {
			resultNode = &ListNode{Val: leastValue}
		} else {
			node := resultNode

			for node.Next != nil {
				node = node.Next
			}

			node.Next = &ListNode{Val: leastValue}
		}

		if lists[leastIdx].Next == nil {
			if len(lists) == 1 {
				return resultNode
			}

			lists[leastIdx] = lists[len(lists)-1]
			lists = lists[:len(lists)-1]
		} else {
			lists[leastIdx] = lists[leastIdx].Next
		}
	}

	return resultNode
}
