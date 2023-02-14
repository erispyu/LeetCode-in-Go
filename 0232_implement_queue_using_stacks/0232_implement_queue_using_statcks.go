package main

import "fmt"

type MyQueue struct {
	popStack  []int
	pushStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			tmp := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
			this.popStack = append(this.popStack, tmp)
		}
	}
	peek := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return peek
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			tmp := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
			this.popStack = append(this.popStack, tmp)
		}
	}
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return (len(this.popStack) + len(this.pushStack)) == 0
}

func main() {
	q := Constructor()
	fmt.Println(q.Empty())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
