# Ordered Map
This package implements an ordered map, using container/list and build-in map, it remembers the original insertion order of the keys, and primarily used in the situation of needing to pay attention to insertion order, for example first in first out.

All operations have O(1) time complexity.

# Example
```
// Create a new map.
m := orderedmap.New()

// Sets item within map, sets 1 under key "a".
m.Set("a", 1)

// Retrieves item from map.
if elem, exist := m.Get("a"); exist {
  val := elem.(int)
  fmt.Println(val)
}

m.Set("b", 2)
m.Set("c", 3)

// Iterates on the map.
for e := m.First; e != nil; e = e.Next() {
  key := item.Key
  val := item.Value.(int)
  fmt.Printf("key: %v, value: %d\n", key, val)
}

// Deletes item under key "a".
m.Delete("a")
```

# License
MIT (see [LICENSE](https://github.com/mazhiyu/orderedmap/blob/master/LICENSE) file)
