package main

import "time"

func main() {
	a := make(chan int, 10)

	go func() {
		a <- 1
		a <- 2
		a <- 3
		close(a)
		println("closed")
	}()

	time.Sleep(time.Second)
	var i int
	var ok bool

	i, ok = <-a // receive new data from closed chan
	println(i, ok)
	i, ok = <-a
	println(i, ok)
	i, ok = <-a
	println(i, ok)
	i, ok = <-a // ok is false
	println(i, ok)
	i, ok = <-a // ok is false
	println(i, ok)
}
