package main

import (
	"fmt"
	"sort"
)

// Sometimes we'll want to sort a collection by something other than its natural order.

type byLength []string

// We implement `sort.Interface` - `Len, Less, Swap` on our type,
// so we can use the `sort` package's generic `Sort` function.
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
