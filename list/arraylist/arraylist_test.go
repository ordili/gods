package arraylist

import (
	"testing"
)

func TestNew(t *testing.T) {
	list1 := New[int]()
	if actualValue := list1.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New[int](1, 2, 3)
	if actualValue := list2.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestAdd(t *testing.T) {
	list1 := New[int]()
	list1.Add(10)
	if actualValue := list1.Size(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, _ := list1.Get(0); actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, 10)
	}
}

func TestIndexOf(t *testing.T) {
	list1 := New[int](0, 1, 2)
	if index := list1.InexOf(0); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}
	if index := list1.InexOf(1); index != 1 {
		t.Errorf("Got %v expected %v", index, 1)
	}
	if index := list1.InexOf(2); index != 2 {
		t.Errorf("Got %v expected %v", index, 2)
	}

	if index := list1.InexOf(100); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}
}
