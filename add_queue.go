package main

import (
	"fmt"
	"github.com/streadway/amqp"
)


// add a rabbitmq queue if it does not exist
func AddQueue(a AmqpHost) {
	url := fmt.Sprint("amqp://", a.GetUid(), ":", a.GetPwd(), "@", a.GetAmqpAddr(), "/", a.GetVhost())
	conn, err := amqp.Dial(url)
	throwErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	throwErr(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		a.GetQueue(),
		false,
		false,
		false,
		false,
		nil,
	)
	throwErr(err, "Failed to instantiate a queue")
	body := a.GetQueue()
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	throwErr(err, "Failed to publish a message")
}
