// Package orderedmap implements an ordered map,
// using container/list and build-in map, it remembers
// the original insertion order of the keys.
//
// All operations have O(1) time complexity.
//
// To iterate over an ordered map (where m is a *OrderedMap):
//	for item := range m.Iter() {
//		key := item.Key
//		value:= item.Value
//		// do something with e.Value
//	}
//
package orderedmap

import (
	"container/list"
)

const chbufSize = 32

// OrderedMap holds key-value pairs and remembers
// the original insertion order of the keys.
//
// `key` stores in the map, map's value is the element of the list.
// for the convenience of iteration, the `Value` of list element stores
// the aggregate data of `key` and `value`.
type OrderedMap struct {
	mapper map[string]*list.Element
	lister *list.List
}

// New return an instance of ordered map.
func New() *OrderedMap {
	return &OrderedMap{
		mapper: make(map[string]*list.Element),
		lister: list.New(),
	}
}

// ElemVal encapsulates the key-value pair.
type ElemVal struct {
	Key   string
	Value interface{}
}

// Set sets the given value under the specified key.
func (m *OrderedMap) Set(key string, value interface{}) {
	elem, exist := m.mapper[key]
	if !exist {
		elem = m.lister.PushBack(&ElemVal{key, value})
		m.mapper[key] = elem
	} else {
		elem.Value.(*ElemVal).Value = value
	}
}

// Get retrieves an values from orderedmap under given key.
// If no value was associated with the given key, will return false.
func (m *OrderedMap) Get(key string) (interface{}, bool) {
	elem, exist := m.mapper[key]
	if !exist {
		return nil, false
	}
	return elem.Value.(*ElemVal).Value, true
}

// Delete deletes an item from the ordered map.
func (m *OrderedMap) Delete(key string) {
	elem, exist := m.mapper[key]
	if !exist {
		return
	}

	m.lister.Remove(elem)
	delete(m.mapper, key)
}

// Len returns the number of elements of ordered map m.
// The complexity is O(1).
func (m *OrderedMap) Len() int { return m.lister.Len() }

// Iter returns a buffered iterator which could be used in a for range loop.
func (m *OrderedMap) Iter() <-chan ElemVal {
	ch := make(chan ElemVal, chbufSize)
	go func() {
		for e := m.lister.Front(); e != nil; e = e.Next() {
			ch <- *e.Value.(*ElemVal)
		}
		close(ch)
	}()
	return ch
}
