package singlyLList

import (
    "fmt"
)

type Node struct {
	Value int
	Next *Node
}


func create (val int) *Node {
	return &Node{val, nil}
}


func add (head *Node, val int) *Node{
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	
	curr.Next = &Node{val, nil}
	return head
}


func printList (head *Node) {
	if head == nil {
		fmt.Printf("empty list")
	}
	
	for head != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next
	}
	fmt.Println()
}


func update (head *Node, old int, new int) *Node {
	curr := head
	for curr != nil {
		if curr.Value == old {
			curr.Value = new
			return head
		}
		curr = curr.Next
	}
	
	fmt.Printf("Element %v not in the linked list\n", old)
	return head	
}


func remove (head *Node, val int) *Node {
	if head.Value == val {
		return head.Next
	}
	
	var prev *Node
	curr := head
	for curr != nil {
		if curr.Value == val {
			prev.Next = curr.Next
			return head
		}
		prev = curr
		curr = curr.Next
	}
	
	fmt.Printf("Element %v not in the linked list\n", val)
	return head	
}


func reverse (head *Node) *Node {
	var prev *Node
	curr, forw := head, head
	
	for curr != nil {
		forw = curr.Next
		curr.Next = prev
		prev = curr
		curr = forw
	}
	return prev
}


func main() {
	m := create(4.00)
	m = add(m, 7)
	m = add(m, 8)
	m = add(m, 9)
	m = add(m, 10)
	printList(m)
	m = update(m, 4, 5)
	m = update(m, 10, 11)
	printList(m)
	m = remove(m, 8)
	m = remove(m, 18)
	printList(m)
	m = reverse(m) 
	printList(m)
	m = remove(m, 11)
	printList(m)
   		


}
