package singlelinkedlist

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func New[T any](vals ...T) *List[T] {
	list := List[T]{nil, nil, 0}
	list.Add(vals...)
	return &list
}

func (list *List[T]) Add(values ...T) {
	for _, val := range values {
		n := node[T]{value: val, next: nil}
		if list.last == nil {
			list.first = &n
			list.last = &n
		} else {
			list.last.next = &n
			list.last = &n
		}
		list.size += 1
	}
}

func (list *List[T]) Remove(index int) bool {
	if !list.withinRange(index) {
		return false
	}

	//remove the first node
	if index == 0 {
		list.first = list.first.next
	}

	cur := 0
	previousNode := list.first
	for cur < index-1 {
		cur += 1
		previousNode = previousNode.next
	}
	if index != 0 {
		previousNode.next = previousNode.next.next
	}

	list.size -= 1
	if list.size == 0 {
		list.first = nil
		list.last = nil
	}
	return true
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) String() string {
	op := make([]string, list.size, list.size)
	n := list.first
	i := 0
	for n != nil {
		op[i] = fmt.Sprintf("%v", n.value)
		i += 1
		n = n.next
	}
	return fmt.Sprintf("lists[%s]", strings.Join(op, ","))
}

// check that the index is in the bounds of the list
func (list *List[T]) withinRange(index int) bool {
	return 0 <= index && index < list.size
}
