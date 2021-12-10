package cqueue

// 剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	appendStack stack
	deleteStack stack
}

func Constructor() CQueue {
	return CQueue{
		appendStack: stack{
			data: []int{},
		},
		deleteStack: stack{
			data: []int{},
		},
	}
}

func (q *CQueue) AppendTail(value int) {
	q.appendStack.Push(value)
}

func (q *CQueue) DeleteHead() int {
	val := q.deleteStack.Pop()
	if val == -1 {
		for {
			val := q.appendStack.Pop()
			if val == -1 {
				break
			}
			q.deleteStack.Push(val)
		}
		val = q.deleteStack.Pop()
	}
	return val
}

type stack struct {
	data []int
}

func (s *stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *stack) Pop() int {
	len := len(s.data)
	if len == 0 {
		return -1
	}
	top := s.data[len-1]
	s.data = s.data[:len-1]
	return top
}
