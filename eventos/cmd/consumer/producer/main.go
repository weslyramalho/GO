package main

import "github.com/weslyramalho/GO/eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "HELLO BROTHER!!", "amq.direct")
}
