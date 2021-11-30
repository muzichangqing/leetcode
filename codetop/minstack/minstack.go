package minstack

type MinStack struct {
	data *stack
	min  *stack
}

func Constructor() MinStack {
	return MinStack{
		data: &stack{
			data: []int{},
			top:  -1,
			size: 0,
		},
		min: &stack{
			data: []int{},
			top:  -1,
			size: 0,
		},
	}
}

func (ms *MinStack) Push(val int) {
	ms.data.push(val)
	if ms.min.isEmpty() || ms.min.getTop() >= val {
		ms.min.push(val)
	}
}

func (ms *MinStack) Pop() {
	val := ms.data.getTop()
	ms.data.pop()
	if !ms.min.isEmpty() && ms.min.getTop() == val {
		ms.min.pop()
	}
}

func (ms *MinStack) Top() int {
	return ms.data.getTop()
}

func (ms *MinStack) GetMin() int {
	return ms.min.getTop()
}

type stack struct {
	data []int
	top  int
	size int
}

func (s *stack) push(val int) {
	s.top++
	if s.top >= s.size {
		s.data = append(s.data, val)
		s.size = len(s.data)
	} else {
		s.data[s.top] = val
	}
}

func (s *stack) pop() {
	s.top--
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) getTop() int {
	return s.data[s.top]
}
