package in

import (
	"github.com/gin-gonic/gin"
	"message-delivery/application/domain"
	"message-delivery/application/port/in"
	"net/http"
)

type EventCollectorHandler struct {
	service in.EventCollector
}

func (h *EventCollectorHandler) CollectEvents(ctx *gin.Context) {
	var message domain.Event
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	err := h.service.Collect(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New message created successfully",
	})
}
