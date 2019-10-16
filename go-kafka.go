package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Message struct {
	Topic       string `json:"topic_name"`
	AggregateID string `json:"aggregate_id"`
	Payload     map[string]interface{}
}

type ReturnMessage struct {
	Topic       string `json:"topic_name"`
	AggregateID string `json:"aggregate_id"`
	Payload     map[string]interface{}
	Status      string `json:"status"`
}

type producedMessage struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

//var producer *kafka.Producer

func main() {
	// Initialize Echo web server
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/eventMessages", messagesPostHandler)
	e.Logger.Fatal(e.Start(":9090"))
}

//func InitKafka() error {
//	var err error
//	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
//	return err
//}
//
//func Produce(topics string, message string) error {
//	deliveryChan := make(chan kafka.Event)
//	err := producer.Produce(&kafka.Message{
//		TopicPartition: kafka.TopicPartition{Topic: &topics, Partition: kafka.PartitionAny},
//		Value:          []byte(message),
//	}, deliveryChan)
//
//	if err != nil {
//		fmt.Printf("Produce failed: %v\n", err)
//	}
//
//	e := <-deliveryChan
//	m := e.(*kafka.Message)
//
//	if m.TopicPartition.Error != nil {
//		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
//	} else {
//		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
//			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
//	}
//
//	close(deliveryChan)
//	return m.TopicPartition.Error
//}

func messagesPostHandler(c echo.Context) (err error) {
	//Retrieve body from http request
	m := new(Message)
	err = c.Bind(m)
	if err != nil {
		return err
	}
	saveMessageToKafka(*m)
	var rtn ReturnMessage
	rtn = ReturnMessage{
		Topic:       m.Topic,
		AggregateID: m.AggregateID,
		Payload:     m.Payload,
		Status:      "success",
	}
	return c.JSON(http.StatusOK, rtn)
}

func saveMessageToKafka(msg Message) {
	fmt.Println("save to kafka")
	jsonString, err := json.Marshal(msg)
	msgStr := string(jsonString)
	fmt.Println(msgStr)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()
	// Produce messages to topic (asynchronously)
	topic := msg.Topic
	for _, word := range []string{string(msgStr)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.AssignedPartitions{}},
			Value:          []byte(word),
		}, nil)
	}
}
