package BelajarGoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Feri susmiyanto"
		fmt.Println("selesai mengirim data di channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeREsponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hallo feri susmiyanto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeREsponse(channel)

	data := <-channel
	fmt.Println(data)
	fmt.Println("selesai")

	time.Sleep(5 * time.Second)
}

// channel kusus mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hi, data channel in"
}

// channel kusus menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	//goroutin mengirim data ke channel
	go OnlyIn(channel)

	//goroutine menerima dan menampilkan data
	go OnlyOut(channel)

	fmt.Println("selesai")

	time.Sleep(5 * time.Second)
}

// Buffered channel digunakan menambah jumlah penampung data yg ada di channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "hello"
		channel <- "feri"
		channel <- "susmiyanto"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

// range channel digunakan untuk menerima dan mengirim data secera terus menerus atau tidak tau kpn berhenti,
//
//	hanya untuk satu channel saja
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data " + data)
	}
	fmt.Println("selesai")
}

// select channel digunakan untuk mengambil atau memilih data dari beberpa channel (lebih dari satu channel)
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeREsponse(channel1)
	go GiveMeREsponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1 " + data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2 " + data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
	fmt.Println("selesai")
}

// Defaul select digunakan apabila di dalam channel belum terdapat data yang dikirimkan,
// agar tidak terjadi deatlock ketika mengambil data, data belum tersedia di channel
func TestDefaulSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeREsponse(channel1)
	go GiveMeREsponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1 " + data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2 " + data)
			counter++
		default:
			fmt.Println("Menunggu data masuk")
		}

		if counter == 2 {
			break
		}
	}
	fmt.Println("selesai")
}
