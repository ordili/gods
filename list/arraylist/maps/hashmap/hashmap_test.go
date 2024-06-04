package hashmap

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	m := New[string, int]()
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected is %v", actualValue, 0)
	}
}

func TestPut(t *testing.T) {
	m := New[string, int]()
	key := "jd"
	val := 100
	m.Put(key, val)
	if actualValue, _ := m.Get(key); actualValue != val {
		t.Errorf("Got %v expected is %v", actualValue, val)
	}
}

func TestGet(t *testing.T) {
	m := New[string, int]()
	key := "jd"
	val := 100
	m.Put(key, val)
	if actualValue, _ := m.Get(key); actualValue != val {
		t.Errorf("Got %v expected is %v", actualValue, val)
	}
}

func TestSize(t *testing.T) {
	m := New[string, int]()
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected is %v", actualValue, 0)
	}

	key := "jd"
	val := 100
	m.Put(key, val)
	if actualValue := m.Size(); actualValue != 1 {
		t.Errorf("Got %v expected is %v", actualValue, 1)
	}

	key = "jd1"
	val = 1001
	m.Put(key, val)
	if actualValue := m.Size(); actualValue != 2 {
		t.Errorf("Got %v expected is %v", actualValue, 2)
	}
}

func TestEmpty(t *testing.T) {
	m := New[string, int]()
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected is %v", actualValue, true)
	}

	key := "jd"
	val := 100
	m.Put(key, val)
	if actualValue := m.Empty(); actualValue != false {
		t.Errorf("Got %v expected is %v", actualValue, false)
	}

}

func TestKeys(t *testing.T) {
	m := New[string, int]()
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		m.Put(key, 1)
	}
	if actualKeys := m.Keys(); !reflect.DeepEqual(actualKeys, keys) {
		t.Errorf("Got %v expected is %v", actualKeys, keys)
	}
}

func TestValues(t *testing.T) {
	m := New[string, int]()
	keys := []string{"a", "b", "c"}
	vals := []int{1, 2, 3}
	for index, key := range keys {
		m.Put(key, vals[index])
	}
	if actualVals := m.Values(); !reflect.DeepEqual(actualVals, vals) {
		t.Errorf("Got %v expected is %v", actualVals, vals)
	}
}
