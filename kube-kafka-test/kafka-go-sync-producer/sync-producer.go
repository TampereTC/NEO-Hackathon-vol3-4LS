package main

import (
	"fmt"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
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

	producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println(err)
		}
		fmt.Println(time.Now())
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Fprintln(os.Stderr, time.Now())

	var enqueued, errors int
	ProducerLoop:
	for i:=0; i<60000;i++ {
		select {
			case producer.Input() <- &sarama.ProducerMessage{Topic: "muntopic12", Key: nil, Value: sarama.StringEncoder("sync message 123")}:
				enqueued++
			case err := <-producer.Errors():
				fmt.Println("Failed to produce message", err)
				errors++
			case <-signals:
				break ProducerLoop
		}
	}

	fmt.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)	
	time.Sleep(1000 * time.Second)
}
