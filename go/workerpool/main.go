package main

import (
	"sync"
	"time"
)

func task() {
	time.Sleep(10 * time.Millisecond)
}

func poolByChan(taskc chan int) {
	t := time.Now()
	q := make(chan struct{}, 500)
	for {
		i, ok := <-taskc
		if !ok {
			break
		}
		q <- struct{}{}
		go func(i int) {
			task()
			<-q
		}(i)
	}
	for len(q) != 0 {
		time.Sleep(1)
	}
	println("chan", time.Now().Sub(t).Seconds())
}

func poolBySync(taskc chan int) {
	t := time.Now()

	wg := &sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			for {
				_, ok := <-taskc
				if !ok {
					wg.Done()
					break
				}
				task()
			}
		}()
	}
	wg.Wait()

	println("sync", time.Now().Sub(t).Seconds())
}
func main() {
	mg := &sync.WaitGroup{}
	mg.Add(1)
	c := make(chan int)
	go func() {
		defer mg.Done()
		poolByChan(c)
	}()
	for i := 0; i < 100000; i++ {
		c <- i
	}
	close(c)
	mg.Wait()

	mg.Add(1)
	c = make(chan int)
	go func() {
		defer mg.Done()
		poolBySync(c)
	}()
	for i := 0; i < 100000; i++ {
		c <- i
	}
	close(c)
	mg.Wait()
}
