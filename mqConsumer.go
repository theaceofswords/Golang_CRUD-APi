package main

import (
	"encoding/json"
	"fmt"
	"golang-training/app/models"
	"golang-training/app/repository"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ErrorMsg(err)
	defer conn.Close()

	ch, err := conn.Channel()
	ErrorMsg(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"queue1", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	ErrorMsg(err)

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	ErrorMsg(err)

	gochan := make(chan bool)

	go func() {
		for d := range msgs {
			var emp models.Employee
			json.Unmarshal(d.Body, &emp)
			repository.CreateEmployee(emp)
			log.Printf("Received a message: %s", emp)
		}
	}()

	log.Printf("Waiting for messages")
	<-gochan

}

func ErrorMsg(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}
