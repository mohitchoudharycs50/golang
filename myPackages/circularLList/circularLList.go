package circularLList

import (
    "fmt"
)

type Node struct {
	Value int
	Next *Node
}

var head, tail *Node

func Create (val int) *Node{
	new := &Node{val, nil}
	head, tail = new, new
	tail.Next = head
	return head
}


// insert at the end
func Insert (val int)  {
	new := &Node{val, nil}
	tail.Next = new
	new.Next = head
	tail = new	
}


//delete from beginning
func Delete () {
	if head == nil {
		fmt.Printf("Empty list\n")
		return
	} else if head == tail {
		head = nil
		tail = nil
		return
	}
	tail.Next = head.Next
	head = head.Next
		
}


func IsCircular (head *Node) {

	if head == nil {
		fmt.Printf("nil pointer passed\n")
		return
	}
	
	curr := head
	for curr.Next != head && curr.Next != nil {
		curr = curr.Next
	}
	
	if curr.Next == head {
		fmt.Printf("List is Circular\n")
	} else {
		fmt.Printf("List is not circular\n")
	}
		
}


func Display () {
	if head == nil {
		fmt.Printf("Empty List\n")
		return
	}
	
	curr := head
	for curr != tail {
		fmt.Printf("%v  ", curr.Value)
		curr = curr.Next
	}
	fmt.Printf("%v\n", curr.Value)
}


func main() {

	Create(1)
	Insert(2)
	Insert(3)
	fmt.Println(head, head.Next, head.Next.Next, head.Next.Next.Next, tail)
 	Display()
	IsCircular(head)
//	k := Node{1, nil}
//	k.Next = &k
//	IsCircular(&k)
//	var l *Node
//	IsCircular(l)
	Delete()
	Display()
	Delete()
	Display()
	Delete()
	Display()	
}
