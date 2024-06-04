package hashmap

import "fmt"

type Map[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}

func (m *Map[K, V]) Put(k K, v V) {
	m.m[k] = v
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	val, found := m.m[k]
	return val, found
}

func (m *Map[K, V]) Remove(k K) {
	delete(m.m, k)
}

func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

func (m *Map[K, V]) Size() int {
	return len(m.m)
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	index := 0
	for key := range m.m {
		keys[index] = key
		index += 1
	}
	return keys
}

func (m *Map[K, V]) Values() []V {

	vs := make([]V, m.Size())
	index := 0
	for _, v := range m.m {
		vs[index] = v
		index += 1
	}
	return vs
}

func (m *Map[K, V]) Clear() {
	clear(m.m)
}

func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
