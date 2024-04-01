package main

import (
	"sync"
)

var mu sync.Mutex

func main() {
	fn := func() {
		mu.Lock()
		defer mu.Unlock()

		println("hello")
	}

	defer fn()
	println("a")

	func() {
		defer func() {
			println("world")
		}()

		println("b")
	}()

	println("c")
}
