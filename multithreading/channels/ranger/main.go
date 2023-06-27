package main

import "fmt"

//thread 1
func main() {
	ch := make(chan int)
	go publich(ch)
	reader(ch)

}

func reader(ch chan int) {
	for x := range ch {
		fmt.Println("reader:", x)
	}
}

func publich(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
