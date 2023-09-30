package main

import (
	"client-go/routes/any"
	chat_completion "client-go/routes/chat-completion"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any(chat_completion.Path, chat_completion.Handler)
	r.NoRoute(any.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1007"
	}
	log.Println("Runng in port:", port)

	err := r.Run(":" + port)
	if err != nil {
		fmt.Println("Gin startup err", err)
	}
}
