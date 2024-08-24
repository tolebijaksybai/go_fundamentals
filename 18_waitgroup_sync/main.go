package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//withoutWait()
	//withWait()

	// Mutex
	//writeWithoutConcurrent()
	//writeWithoutMutex()

	readWithMutex()
	readWithRWMutex()
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mx      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mx.RLock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mx.RUnlock()
		}()

		go func() {
			defer wg.Done()

			mx.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mx.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mx      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mx.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mx.Unlock()
		}()

		go func() {
			defer wg.Done()

			mx.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mx.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(time.Nanosecond)

			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func withWait() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		//wg.Add(1) // тоже можно

		go func(i int) {

			fmt.Println(i + 1)

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Exit")
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}
