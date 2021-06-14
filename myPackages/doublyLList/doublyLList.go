package doublyLList

import (
    "fmt"
)

type Node struct {
	Value int
	Next *Node
	Prev *Node
}


func create (val int) *Node {
	return &Node{val, nil, nil}
}


func add (head *Node, val int) *Node{
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	
	curr.Next = &Node{val, nil, nil}
	curr.Next.Prev = curr
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
	
	fmt.Printf("Element not in the linked list\n")
	return head	
}


func remove (head *Node, val int) *Node {
	if head.Value == val {
		head = head.Next
		head.Prev = nil
		return head
	}
	
	curr := head
	for curr.Next != nil {
		if curr.Value == val {
			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev
			return head
		}
		curr = curr.Next
	}
	
	if curr.Value == val {
		curr.Prev.Next = nil
	} else {	
		fmt.Printf("Element %v not in the linked list\n", val)
	}
	return head	
}


func reverse (head *Node) *Node {
	var next *Node
	curr, forw := head, head
	switch {
	case head == nil :
		fmt.Printf("Nil pointer dereference\n")
		return head
	
	case head.Prev != nil:
		fmt.Printf("Argument passed is not the head of the list\n")
		return head
		
	case head.Prev == nil && head.Next == nil:
		return head	
	}
	
	for curr != nil {
		forw = curr.Next
		curr.Prev = curr.Next
		curr.Next = next
		next = curr
		curr = forw		
	}
	return next
}


func main() {
	m := create(4)
	fmt.Println(m)
	m = add(m, 7)
	m = add(m, 8)
	m = add(m, 9)
	m = add(m, 10)
	m = update(m, 4, 5)
	printList(m)
	m = remove(m, 15)
	m = remove(m, 5)
	printList(m)
	m = add(m, 8)
	printList(m)
	m = remove(m, 8)
	printList(m)
	p := Node{}
	var k *Node = &p
	k = reverse(k.Next) 
	printList(k)
/**

   	for m != nil {
		fmt.Println(*m)
		m = m.Next
	}
*/	

}
