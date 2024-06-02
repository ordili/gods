package arraylist

const growthFactor = float32(1.5)

type List[T comparable] struct {
	elements []T
	size     int
}

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, val := range values {
		list.elements[list.size] = val
		list.size += 1
	}
}

func (list *List[T]) Get(index int) (T, bool) {
	if !list.withinRange(index) {
		var ret T
		return ret, false
	}
	return list.elements[index], true
}

func (list *List[T]) Remove(index int) (T, bool) {
	if !list.withinRange(index) {
		var ret T
		return ret, false
	}

	ret := list.elements[index]
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size -= 1
	return ret, true
}

func (list *List[T]) withinRange(index int) bool {

	return 0 <= index && index < list.size
}

func (list *List[T]) growBy(delt int) {

	currentCapacity := cap(list.elements)
	if len(list.elements)+delt >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+delt))
		list.reSize(newCapacity)
	}
}

func (list *List[T]) reSize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Contains(target T) bool {
	for _, val := range list.elements {
		if target == val {
			return true
		}
	}
	return false
}
