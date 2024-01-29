/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    slow := head
    fast := head

    for fast!=nil && fast.Next!=nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    if fast!=nil && fast.Next == nil{
        slow = slow.Next
    }

    var prev *ListNode //NULL pointer
    var temp *ListNode

    for slow != nil && slow.Next != nil{
        temp = slow.Next
        slow.Next = prev
        prev = slow
        slow = temp
    }

    if slow != nil{
        slow.Next = prev
    }

    fast = head
    for slow != nil && fast != nil{
        if slow.Val != fast.Val{
            return false
        }
        slow = slow.Next
        fast = fast.Next
    }
    
    return true

}