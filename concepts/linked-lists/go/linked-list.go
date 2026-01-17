package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 1. Create a new list
	l := list.New()

	// 2. Add items (Push)
	// Front <-> ["Apple"] <-> Back
	e := l.PushFront("Apple")

	// Front <-> ["Banana"] <-> ["Apple"] <-> Back
	l.PushFront("Banana")

	// Front <-> ["Banana"] <-> ["Apple"] <-> ["Cherry"] <-> Back
	l.PushBack("Cherry")

	printList("Initial State", l)

	// 3. Move to Front
	// Let's say we accessed "Apple" (e1). We move it to front.
	l.MoveToFront(e)
	printList("After Accessing Apple (MoveToFront)", l)

	// 4. Eviction
	// To remove the last item, we first get the Element...
	lastElement := l.Back()
	if lastElement != nil {
		// ...then we remove it.
		// Note: Remove returns the value that was stored.
		val := l.Remove(lastElement)
		fmt.Printf("Evicted: %v\n", val)
	}
	printList("After Eviction", l)
}

func printList(title string, l *list.List) {
	fmt.Printf("--- %s ---\n", title)
	// Loop from Front to Back
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%v] -> ", e.Value)
	}
	fmt.Println("nil")
}
