package main

import (
	"context"
	"fmt"
	"sync"
)

type Lenther map[rune]int

var buffer = make(chan string, 1000)

func do(arr []string, concurrentCount uint) Lenther {
	ch := make(chan string, concurrentCount)
	result := make(map[rune]int)
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	// consumer
	go func() {
		for {
			str := <-ch
			go func() {
				defer wg.Done()
				rw.Lock()
				for i := range str {
					result[rune(str[i])]++
				}
				rw.Unlock()
			}()
		}
	}()
	for i := range arr {
		buffer <- arr[i]
		wg.Add(1)
	}
	close(buffer)
	go producer(ch)

	wg.Wait()
	return result
}

func producer(ch chan string) {
	for {
		str, ok := <-buffer
		if !ok {
			break
		}
		ch <- str
	}
}

func main() {
	arr := []string{"abc", "213ab", "cdsadfa"}
	res := do(arr, 2)
	ctx:= context.WithCancel()
	fmt.Println(res)
}
