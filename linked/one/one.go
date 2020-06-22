/**
 * Created by Adwind.
 * User: liuyunlong
 * Date: 6/22/20
 * Time: 23:04
 */
package main

import "fmt"

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

//创建链表
func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

//question 如何实现链表的逆序

func listReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur

	}
	node.Next = pre

}

func main() {
	head := &LNode{}
	fmt.Println("就地逆序")
	CreateNode(head, 8)
	PrintNode("逆序前", head)
	listReverse(head)
	PrintNode("逆序后", head)
}
