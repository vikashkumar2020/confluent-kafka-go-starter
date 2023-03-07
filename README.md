# confluent-kafka-go-starter

## Description

This is kafka implementation in golang using confluent-kafka package written in go. Bootstraped code to make it simpler to use and have producer consumer pattern. 
Multiple consumer reading messages from queue.

## Requirements 

- docker 
- go

## Set Up

Run docker compose file to have docker container for kafka broker and zookeeper.

Use Command - docker compose up

Then run main.go file for every section.

Producer main.go file produce messages
Consumer1 and Consumer2 read from kafka message queue of particular topic indipendently.



