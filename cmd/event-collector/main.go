package main

import (
	"github.com/gin-gonic/gin"
	"message-delivery/adapter/in"
	"message-delivery/application/service"
)

func main() {
	svc := service.NewEventCollectorService()
	handler := in.NewEventCollectorHandler(svc)

	engine := gin.New()
	engine.POST("/api/v1/events", func(context *gin.Context) {
		handler.CollectEvents(context)
	})
	err := engine.Run(":3000")
	if err != nil {
		panic(err)
	}
}
