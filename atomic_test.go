package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}
	
	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}

// var group sync.WaitGroup
// var counter int64 = 0
// for i := 0; i < 100; i++ {
// 	group.Add(1)
// 	go func() {
// 		for j := 0; j < 100; j++ {
// 			atomic.AddInt64(&counter, 1)
// 		}
// 		group.Done()
// 	}()
// }

// group.Wait()
// fmt.Println("Counter", counter)
