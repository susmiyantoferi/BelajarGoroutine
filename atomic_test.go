package BelajarGoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Atomic digunakan untuk menggunakan tipe data primitive secara aman pada proses concurrent
// Atomic aman dari race condition

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)

			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
				//x = x + 1
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("hasil x = ", x)
}
