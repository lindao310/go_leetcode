package valid_parentheses

//先得实现个栈
import "sync"

type (
	node struct {
		value interface{}
		prev *node
	}
	Stack struct {
		top  *node
		length int
		lock sync.RWMutex
	}
)
//create stack
func NewStack() *Stack {
	return &Stack{nil, 0, sync.RWMutex{}}
}
//return the number of items in the stack
func (this *Stack)len() int {
	return this.length
}

//view the top item value
func (this *Stack) peak() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

//pop the top item and return it
func (this *Stack) pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

//push item to stack
func (this *Stack) push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()

	n := &node{value, this.top}
	this.top = n

	this.length++
}
