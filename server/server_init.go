package server

import (
	"log"
	h "task/handlers"

	"github.com/gin-gonic/gin"
)

func ServerInit(h *h.HandlerModule) {
	r := gin.Default()
	r.POST("/numbers", h.CreateNumberHandler)
	if err := r.Run(); err != nil {
		log.Println("Server init error", err)
		return
	}
}
