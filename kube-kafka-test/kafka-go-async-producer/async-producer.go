package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("kafka:9092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
)

func main() {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			fmt.Println(err)
			errors++
		}
	}()

	fmt.Fprintln(os.Stderr, time.Now())

ProducerLoop:
	for i:=0 ; i< 60000 ; i++ {
		message := &sarama.ProducerMessage{Topic: "muntopic12", Value: sarama.StringEncoder("async message 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	fmt.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
	time.Sleep(1000 * time.Second)
}
