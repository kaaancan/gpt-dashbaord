package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func handler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	req, err := http.NewRequest(c.Request.Method, "https://api.openai.com/v1"+c.Request.URL.Path, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Authorization", c.GetHeader("Authorization"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func main() {
	r := gin.Default()
	r.Any("/*any", handler)
	err := r.Run(":3000")
	if err != nil {
		fmt.Println("Gin startup err", err)
	}
}
