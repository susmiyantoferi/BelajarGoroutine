package BelajarGoroutine

import (
	"fmt"
	"sync"
	"testing"
)

//sync.Pool secara simple nya digunakan untuk menyimpan data, selanjutnya kita bisa menggunakan datanya,
// jika sudah tidak digunakan datanya kita bisa mengembalikan datanya ke dalam variable pool

func TestPool(t *testing.T) {
	pool := sync.Pool{

		//default jika data di dalam pool == nill
		New: func() interface{} {
			return "New"
		},
	}

	group := sync.WaitGroup{}

	//menambah data ke dalam pool
	pool.Put("Hello")
	pool.Put("Feri")
	pool.Put("Susmiyanto")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)

			//mengabil data dari pool
			data := pool.Get()

			fmt.Println(data)

			//mengembalikan data ke dalam pool
			pool.Put(data)
			group.Done()
		}()
		group.Wait()
	}
	fmt.Println("selesai")
}
