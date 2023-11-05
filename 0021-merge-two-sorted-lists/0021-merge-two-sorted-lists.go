/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }

    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    var resultNode *ListNode
    var currentNode *ListNode

    for list1 != nil || list2 != nil {
        var currentVal int

        if list1 == nil {
            currentVal = list2.Val
            list2 = list2.Next
        } else if list2 == nil {
            currentVal = list1.Val
            list1 = list1.Next
        } else if list1.Val < list2.Val {
            currentVal = list1.Val
            list1 = list1.Next
        } else {
            currentVal = list2.Val
            list2 = list2.Next
        }

        if resultNode == nil {
            resultNode = &ListNode{Val: currentVal}
            currentNode = resultNode
        } else {
            currentNode.Next = &ListNode{Val: currentVal}
            currentNode = currentNode.Next
        }
    }

    return resultNode
}