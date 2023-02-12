package top100

type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) move() {
	for len(this.pushStack) != 0 {
		this.popStack = append(this.popStack, this.pushStack[len(this.pushStack)-1])
		this.pushStack = this.pushStack[:len(this.pushStack)-1]
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	v := this.Peek()
	this.popStack = this.popStack[:len(this.popStack)-1]
	return v
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		this.move()
	}
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.popStack) == 0 && len(this.pushStack) == 0
}
