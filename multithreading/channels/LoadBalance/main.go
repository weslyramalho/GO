package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("wORK %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
func main() {
	data := make(chan int)
	QtWorkers := 200

	//inicaliza os workers
	for i := 1; i <= QtWorkers; i++ {
		go worker(i, data) //chama o metodo de forma anonima
	}

	for i := 0; i < 1000; i++ {
		data <- i
	}
}
