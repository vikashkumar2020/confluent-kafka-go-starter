package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)



func main(){
	topic := "HVSE"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":         "consumer2",
		"auto.offset.reset":  "smallest",
	})

	if err !=nil{
		log.Fatal(err) 
	}
	err = consumer.Subscribe(topic,nil)

	if err!=nil{
		log.Fatal(err)
	}

	for{
		ev := consumer.Poll(10)

		switch e := ev.(type){
		case *kafka.Message:
			fmt.Printf("consumed message by consumer 2 : %s\n", string(e.Value))

		case *kafka.Error:
			fmt.Printf("%v\n",e)
		}
	}
}