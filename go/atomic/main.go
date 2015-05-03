package main

import (
	"sync"
	"sync/atomic"
)

const max = 1000000

func main() {
	var a int32
	a = 0

	println("expected:", max*2)
	wg := &sync.WaitGroup{}
	f := func() {
		for i := 0; i < max; i++ {
			a++
		}
		wg.Done()
	}

	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	println(a)

	a = 0
	f = func() {
		for i := 0; i < max; i++ {
			atomic.AddInt32(&a, 1)
		}
		wg.Done()
	}
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	println(a)
}
