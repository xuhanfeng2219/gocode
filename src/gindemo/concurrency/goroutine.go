package concurrency

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("hello goroutine!")
}

func main() {
	hello()
	fmt.Println("main goroutine done!")
}

func main1() {
	go hello()
	fmt.Println("ll")
}

var wg sync.WaitGroup

func hello1() {
	defer wg.Done()
	fmt.Println("group hello")
}

func main2() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello1()
	}
	wg.Wait()
}
