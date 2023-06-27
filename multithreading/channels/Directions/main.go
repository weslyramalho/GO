package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}
func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}
