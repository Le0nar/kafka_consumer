package main

import (
	"github.com/Le0nar/kafka_consumer/internal/handler"
	"github.com/Le0nar/kafka_consumer/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run(":8081")
}
