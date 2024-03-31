package BelajarGoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

//gomaxprocs digunakan untuk melihat berapa banyak thread yang berjalan dan mengubah jumlah thread
//secara default jumlah thread di golang sebanyak jumlah jumlah CPU komputer

func TestGomaxprocs(t *testing.T) {
	//mengambil total CPU yang ada
	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu: ", totalCpu)

	//mengambil total thread yg berjalan
	//-1 digunakan untuk mengambil total thread yg ada
	//apabila n = diatas 0, digunakan mengubah jumlah thread yg ada
	//runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread yg berjalan: ", totalThread)

	//mengambil total goroutine yg menyala
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("goroutine yg menyala: ", totalGoroutine)
}
