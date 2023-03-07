package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// entity to create message
type OrderPlacer struct {
	producer   *kafka.Producer  // producer Config
	topic      string           // topic name
	deliverych chan kafka.Event // message delivery confirmation channel
}

// factory method to create messages
func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 1000),
	}
}

// message Producer
func (op *OrderPlacer) OrderPlacer(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	},
		op.deliverych,
	)

	if err != nil {
		log.Fatal(err)
	}

	<-op.deliverych

	fmt.Printf("MSG produced %s\n", format)
	return nil;
}

func main() {

	topic := "HVSE"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "foo",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	op := NewOrderPlacer(p, topic)

	for i := 0; i < 1000; i++ {
		if err := op.OrderPlacer("market", i+1); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 1) // to slow down producing messages
	}

}
