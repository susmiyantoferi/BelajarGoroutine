package BelajarGoroutine

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map aman dari race condition
// sync.Map mirip dengan tipe data Map biasa,namun yg berbeda sync.map aman untuk penggunaan concurrent di goroutine
// Store(key,value) : menyimpan data ke sync.Map menggunakan key
// Load(key) : mengambil data dari sync.Map menggunakan key
// Delete(key) : menghapus data dari sync.Map menggunakan key
// Range(func(key,value)) : digunakan untuk melakukan iterasi seluruh data di sync.Map
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ": ", value)
		return true
	})
}
