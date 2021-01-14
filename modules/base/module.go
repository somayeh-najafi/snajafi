package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func declareResponseQueue(ch *amqp.Channel, queue string) amqp.Queue {
	return declareQueue(ch, fmt.Sprintf("%v-response", queue))
}

func declareQueue(ch *amqp.Channel, queue string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}

func main() {

	// Connecting to RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_ADDRESS"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Retrieving the correct queue
	queueName := os.Getenv("QUEUE_NAME")
	log.Printf("Listening to queue: %s", queueName)
	queue := declareQueue(ch, queueName)
	responseQueue := declareResponseQueue(ch, queueName)

	// Consuming a message
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	// Sending back the module result
	responded := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Dummy computation
			dummyComputationTimeInSeconds := 15
			log.Printf("Performing dummy computation for %d seconds...", dummyComputationTimeInSeconds)
			time.Sleep(time.Duration(dummyComputationTimeInSeconds) * time.Second)

			// Sending a response
			log.Printf("...dummy computation completed. Sending a response to the master...")
			responseBody := fmt.Sprintf("Response from %v", queueName)
			err = ch.Publish(
				"",        // exchange
				responseQueue.Name, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: queueName,
					Body:          []byte(responseBody),
				})
			failOnError(err, "Failed to publish the response")
			log.Printf("...response sent.")
			responded <- true
		}
	}()
	<-responded
}
