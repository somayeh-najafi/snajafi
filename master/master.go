package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
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

func publishMessage(ch *amqp.Channel, q *amqp.Queue, body string) {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			ContentType:   "text/plain",
			//CorrelationId: q.Name,
			//ReplyTo:       fmt.Sprintf("%v-response", q.Name),
			Body:          []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func consumeModuleResponse(ch *amqp.Channel, q *amqp.Queue, finalModuleHasResponded chan <- bool, remainingModules *int) {
	moduleResponse, err := ch.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, fmt.Sprintf("Failed to consume the %v response", q.Name))

	// Printing to screen
	go func() {
		for d := range moduleResponse {
			log.Printf("Received a message from a %v: %s", q.Name, d.Body)
		}
		*remainingModules -= 1
		if *remainingModules == 0 {
			finalModuleHasResponded <- true
		}
	}()
}

func main() {

	// Connecting to RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_ADDRESS"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaring queues
	builderQueue := declareQueue(ch, "builder")
	builderResponseQueue := declareResponseQueue(ch, "builder")
	analyzerQueue := declareQueue(ch, "analyzer")
	analyzerResponseQueue := declareResponseQueue(ch, "analyzer")
	reporterQueue := declareQueue(ch, "reporter")
	reporterResponseQueue := declareResponseQueue(ch, "reporter")

	type module struct {
		name string
		queue amqp.Queue
	}

	// Modules who are going to be waked up, in order
	modulesToWakeUp := []module{
		{
			name: "builder",
			queue: builderQueue,
		},
		{
			name: "analyzer",
			queue: analyzerQueue,
		},
		{
			name: "reporter",
			queue: reporterQueue,
		},
	}

	// Waking up one Pod of each module at the time
	for _, module := range modulesToWakeUp {
		sleepTimeInSeconds := 5
		time.Sleep(time.Duration(sleepTimeInSeconds) * time.Second)
		log.Printf("Waking up one %s...", module.name)
		publishMessage(ch, &module.queue, fmt.Sprintf("Go, %s!", module.name))
	}

	// Waiting for modules' responses.
	// Apparently, we cannot wait for multiple messages belonging to different queues.
	// However, it makes sense for Quartermaster to wait only for messages belonging
	// to the current phase (i.e., builder, analysis, report).
	// The master consumes messages following that order.
	allModulesHaveResponded := make(chan bool)
	remainingModules := len(modulesToWakeUp)
	consumeModuleResponse(ch, &builderResponseQueue, allModulesHaveResponded, &remainingModules)
	consumeModuleResponse(ch, &analyzerResponseQueue, allModulesHaveResponded, &remainingModules)
	consumeModuleResponse(ch, &reporterResponseQueue, allModulesHaveResponded, &remainingModules)

	<-allModulesHaveResponded
}
