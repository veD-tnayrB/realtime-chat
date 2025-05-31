package main

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/veD-tnayrB/chat/cmd/server/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port not defined")
	}

	hub := models.Hub{
		Conns:         make(models.Conns),
		Events:        make(chan models.Event),
		Mutx:          sync.Mutex{},
		Subscriptions: make(map[string]models.Conns),
	}
	defer hub.Close()

	go hub.HandleEvents()

	router := gin.Default()
	ws := router.Group("/ws")
	ws.GET("/connect", hub.Connect)
	router.Run(port)
}
