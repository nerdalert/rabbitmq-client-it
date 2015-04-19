package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

// delete a rabbitmq queue
func DeleteQueue(a AmqpHost) {
	url := fmt.Sprint("amqp://", a.GetUid(), ":", a.GetPwd(), "@", a.GetAmqpAddr(), "/", a.GetVhost())
	conn, err := amqp.Dial(url)
	throwErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	throwErr(err, "Failed to open a channel")
	defer ch.Close()
	delQueue, err := ch.QueueDelete(
		a.GetQueue(),
		false,
		false,
		false,
	)
	if delQueue == 0 {
		log.Fatalf("Deleting queue [ %s ] failed, try testcase #1 first to verify the queue existed.", a.GetQueue())
	} else {
		log.Infof("Deletion of [ %s ] appeared to be successful.", a.GetQueue())
	}
}
