package BelajarGoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.cond digunakan sebagai implementasi locking berbasis kondisi
// sync.cond membutuhkan locking diantaranya sync.mutex atau sync.rwmutex
// pada sync.cond terdapat function Wait() yang digunakan apakah perlu menunggu atau tidak
// pada sync.cond terdapat function Signal() yang digunakan agar satu persatu goroutine tidak perlu menunggu lagi
// pada sync.cond terdapat function Broadcast() yang digunakan untuk memberi tahu semua goroutine agar
// tidk perlu menunggu lagi
var group = sync.WaitGroup{}
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Selesai", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)

	}

	go func() {
		for a := 0; a < 10; a++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
			//cond.Broadcast()
		}
	}()

	group.Wait()
}
