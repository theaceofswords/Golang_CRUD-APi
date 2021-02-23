package controllers

import (
	"encoding/json"
	"fmt"
	"golang-training/app/models"
	"log"

	"github.com/streadway/amqp"
)

func MQProducer(emp models.Employee) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	defer conn.Close()

	fmt.Println(" Connected to RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"queue1", // name
		false,    // durable
		false,    // autodelete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	body, _ := json.Marshal(emp)
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "json",
			Body:        []byte(body),
		})
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	log.Printf(" [x] Sent %s", body)
}
