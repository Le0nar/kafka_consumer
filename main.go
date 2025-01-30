package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Le0nar/kafka_consumer/internal/handler"
	"github.com/Le0nar/kafka_consumer/internal/order"
	"github.com/Le0nar/kafka_consumer/internal/service"
	"github.com/segmentio/kafka-go"
)

func readFromKafka() {
	// Создаем Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "consumer-group-1", // Используем уникальную группу
	})

	defer reader.Close()

	for {
		// Читаем сообщение
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}

		var order order.Order
		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Fatal("Error unmarshaling message:", err)
		}

		// Выводим сообщение в консоль
		fmt.Println("Consumer 1 received order:", order)
	}
}

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	go readFromKafka()

	router.Run(":8081")
}
