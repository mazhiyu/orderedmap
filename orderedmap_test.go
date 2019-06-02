package orderedmap

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMapCreation(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("map is null.")
	}

	if m.Len() != 0 {
		t.Error("new map should be empty.")
	}
}

func TestSet(t *testing.T) {
	m := New()

	m.Set("a", 1)
	m.Set("b", 2)

	if m.Len() != 2 {
		t.Error("map should contain exactly two elements.")
	}
}

func TestGet(t *testing.T) {
	m := New()

	// Get a missing element.
	val, exist := m.Get("a")

	if exist {
		t.Error("exist should be false when item is missing from map.")
	}

	if val != nil {
		t.Error("Missing values should return as null.")
	}

	m.Set("a", 1)
	m.Set("a", 2)

	if m.Len() != 1 {
		t.Error("map should contain exactly one element.")
	}

	// Retrieve inserted element.
	elem, exist := m.Get("a")
	elemval := elem.(int) // Type assertion.

	if !exist {
		t.Error("exist should be true for item stored within the map.")
	}

	if &elemval == nil {
		t.Error("expecting a number, not null.")
	}

	if elemval != 2 {
		t.Error("item was modified.")
	}
}

func TestDelete(t *testing.T) {
	m := New()

	m.Set("a", 1)

	m.Delete("a")

	if m.Len() != 0 {
		t.Error("Expecting count to be zero once item was removed.")
	}

	temp, exist := m.Get("a")

	if exist != false {
		t.Error("Expecting exist to be false for missing items.")
	}

	if temp != nil {
		t.Error("Expecting item to be nil after its removal.")
	}

	// Remove a none existing element.
	m.Delete("noone")
}

func TestOrderedMap(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("map is null.")
	}

	if m.Len() != 0 {
		t.Error("new map should be empty.")
	}

	m.Set("a", 1)
	fmt.Println()
}

func TestLen(t *testing.T) {
	m := New()
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	if m.Len() != 100 {
		t.Error("Expecting 100 element within map.")
	}
}

func TestIterator(t *testing.T) {
	m := New()

	// Insert 100 elements.
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), i)
	}

	counter := 0
	// Iterate over elements.
	for item := range m.Iter() {
		key := item.Key
		val := item.Value

		if key != strconv.Itoa(counter) {
			t.Error("key was modified.")
		}

		if val.(int) != counter {
			t.Error("val was modified.")
		}

		if val == nil {
			t.Error("Expecting an object.")
		}
		counter++
	}

	if counter != 100 {
		t.Error("We should have counted 100 elements.")
	}
}
