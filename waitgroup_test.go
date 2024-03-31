package BelajarGoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Waitgroup digunakan untuk menunggu proses goroutine ataupun proses lain sampai selesai,
// sebelum applikasi selesai berjalan
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestRunAsynchronous(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("selesai")
}

// sync.Once digunakan hanya untuk menjalankan goroutine sekali saja, hanya goroutine yg pertama yang dapat perjalan
// yang lain tidak akan bisa berjalan
var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter", counter)
}
