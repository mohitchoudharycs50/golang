package queue

import (
	"fmt"
)


func Enqueue (queue []int, val int) []int {
	queue = append(queue, val)
	return queue
}


func Dequeue (queue []int) []int {
	return queue[1:]
}


func IsEmpty (queue []int) {
	if len(queue) == 0 {
		fmt.Printf("Size of the queue is 0\n")
	} else {
		fmt.Println(len(queue) == 0)
	}
}


func IsFull (queue []int) {
	fmt.Println(len(queue) == cap(queue))
}


func Peek (queue []int) int {
	return queue[0]
}

func main() {
	q := make([]int, 0)
	IsEmpty(q)
	IsFull(q)
	q = Enqueue(q, 1)
	fmt.Println(len(q), cap(q))
	q = Enqueue(q, 2)
	fmt.Println(len(q), cap(q))
	q = Enqueue(q, 3)
	fmt.Println(len(q), cap(q))
	q = Enqueue(q, 4)
	fmt.Println(len(q), cap(q))
	IsEmpty(q)
	IsFull(q)
	q = Enqueue(q, 5)
	fmt.Println(len(q), cap(q))
	q = Dequeue(q)
	fmt.Println(len(q), cap(q))
	q = Dequeue(q)
	fmt.Println(len(q), cap(q))
	q = Enqueue(q, 1)
	q = Enqueue(q, 2)
	q = Enqueue(q, 6)
	q = Enqueue(q, 7)
	fmt.Println(q)
	fmt.Println(len(q), cap(q))
	fmt.Println(Peek(q))
}

