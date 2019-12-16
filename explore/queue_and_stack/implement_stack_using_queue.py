class MyStack:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.__mq = []
        
    def push(self, x: int) -> None:
        """
        Push element x onto stack.
        """
        self.__mq.append(x)

    def pop(self) -> int:
        """
        Removes the element on top of the stack and returns that element.
        """
        sq = []
        for index in range(0, len(self.__mq) - 1):
            sq.append(self.__mq[index])
        element = self.__mq[len(self.__mq) - 1]
        self.__mq = sq
        return element

    def top(self) -> int:
        """
        Get the top element.
        """
        sq = []
        for index in range(0, len(self.__mq) - 1):
            sq.append(self.__mq[index])
        element = self.__mq[len(self.__mq) - 1]
        self.__mq = sq
        self.__mq.append(element)
        return element

    def empty(self) -> bool:
        """
        Returns whether the stack is empty.
        """
        return len(self.__mq) == 0

if __name__ == '__main__':
    obj = MyStack()
    obj.push(1)
    obj.push(2)
    print(obj.pop())
    print(obj.top())
    print(obj.empty())

# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()