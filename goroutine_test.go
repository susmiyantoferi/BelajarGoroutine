package BelajarGoroutine

import (
	"fmt"
	"testing"
	"time"
)

func Hello() {
	fmt.Println("hello world")
}

func TestHello(t *testing.T) {
	go Hello()
	fmt.Println("hai")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("number", number)
}

func TestDIsplay(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
