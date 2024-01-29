type MyQueue struct {
	queue list.List
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyQueue) Pop() int {
	element := this.queue.Front()
	this.queue.Remove(element)
	return element.Value.(int)
}

func (this *MyQueue) Peek() int {
	return this.queue.Front().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.queue.Len() == 0

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */