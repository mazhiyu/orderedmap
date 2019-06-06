// Package orderedmap implements an ordered map,
// using container/list and build-in map, it remembers
// the original insertion order of the keys.
//
// All operations have O(1) time complexity.
//
// To iterate over an ordered map (where m is a *OrderedMap):
//	for e:= m.Front; e != nil; e = e.Next() {
//		key := e.Key
//		value:= e.Value
//		// do something with e.Value
//	}
//
// If you want to delete element while iterating,
// MUST use the following pattern:
//  var next *orderedmap.Element
// 	for e := m.First(); e != nil; e = next {
//		key := e.Key
// 		// assign e.Next() to the next before deleting e
//  	next = e.Next()
// 		m.Delete(key)
// 	}
package orderedmap

import (
	"container/list"
)

// const chbufSize = 32

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

// ElemVal encapsulates the key-value pair which
// stores in the `Value` field of list element.
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

// Get retrieves an values from ordered map under given key.
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
//
// Deprecated: using channel as iterator, you can't break in a `for` loop,
// otherwise the following goroutine will block forever and cause goroutine leak.
// func (m *OrderedMap) Iter() <-chan ElemVal {
// 	ch := make(chan ElemVal, chbufSize)
// 	go func() {
// 		var next *list.Element
// 		for e := m.lister.Front(); e != nil; e = next {
// 			// assign e.Next() to `next` before sending `e.Value` to channel,
// 			// make the delete operation while iterating safe.
// 			next = e.Next()
// 			ch <- *e.Value.(*ElemVal)
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// Element encapsulates the underlying list element and the map's
// key-value pair, to provide a `Next` method for iterating.
type Element struct {
	elem *list.Element
	*ElemVal
}

// First returns the first element of ordered map or nil if the orded map is empty.
func (m *OrderedMap) First() *Element {
	front := m.lister.Front()

	if front == nil {
		return nil
	}

	return &Element{
		elem: front,
		ElemVal: &ElemVal{
			Key:   front.Value.(*ElemVal).Key,
			Value: front.Value.(*ElemVal).Value,
		},
	}
}

// Next returns the next ordered map element or nil.
func (e *Element) Next() *Element {
	next := e.elem.Next()

	if next == nil {
		return nil
	}

	return &Element{
		elem: next,
		ElemVal: &ElemVal{
			Key:   next.Value.(*ElemVal).Key,
			Value: next.Value.(*ElemVal).Value,
		},
	}
}
