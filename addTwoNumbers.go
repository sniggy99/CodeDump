// Problem statement:
//You are given two non-empty linked lists representing two non-negative integers. 
//The digits are stored in reverse order, and each of their nodes contains a single digit. 
//Add the two numbers and return the sum as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
    total := 0
    carry := 0
    list3 := &ListNode{}
    node := list3

    for list1 != nil || list2 != nil {
        total = 0

        if list1 != nil {
            total += list1.Val
            list1 = list1.Next
        }

        if list2 != nil {
            total += list2.Val
            list2 = list2.Next
        }
        total += carry
        carry = total / 10
        total %= 10
        list3.Next = &ListNode{Val: total}
        list3 = list3.Next
    }

    if carry != 0 {
        list3.Next = &ListNode{Val: carry}
    }

    return node.Next
}
