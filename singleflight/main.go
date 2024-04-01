package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

var (
	sg singleflight.Group
	wg sync.WaitGroup
)

func main() {

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err, shared := sg.Do("key", fromData)
			if err != nil {
				panic(err)
			}

			fmt.Printf("v: %v, shared: %v\n", v, shared)
		}()
	}
	wg.Wait()
}

func fromData() (any, error) {
	println("entry....")
	return "hello", nil
}
