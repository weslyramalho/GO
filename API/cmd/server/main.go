package main

import "github.com/weslyramalho/GO/tree/main/API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
