package main

import "fmt"

//Thred 1
func main() {
	canal := make(chan string) //vazio

	//tread2
	go func() {
		canal <- "ola mundo" // esta cheia
	}()

	//tread 1
	msg := <-canal // canal vazio
	fmt.Println(msg)
}
