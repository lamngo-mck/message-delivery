package main

import (
	"github.com/gin-gonic/gin"
	"message-delivery/adapter/in"
	"message-delivery/application/service"
)

func main() {
	svc := service.NewEventCollectorService()
	validator := service.EventValidator{}
	handler := in.NewEventCollectorHandler(svc, validator)

	engine := gin.New()
	engine.POST("/api/v1/events", handler.CollectEvents)
	err := engine.Run(":3000")
	if err != nil {
		panic(err)
	}
}
