/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNode(0, l1, l2)
}

func addNode(co int, l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        if co == 1 {
            return &ListNode{Val: 1}
        }
        
        return nil
    }

    resNode := &ListNode{}
	res := co
    
    var l1Next *ListNode
    var l2Next *ListNode

	if l1 != nil {
		res += l1.Val
        l1Next = l1.Next
	}
    
    if l2 != nil {
		res += l2.Val
        l2Next = l2.Next
	}

	if res > 9 {
        resNode.Val = res - 10
        resNode.Next = addNode(1, l1Next, l2Next)
	} else {
		resNode.Val = res
        resNode.Next = addNode(0, l1Next, l2Next)
	}

    return resNode
}