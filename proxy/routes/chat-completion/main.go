package chat_completion

import (
	"bytes"
	"client-go/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
)

const Path = "/chat/completions"

func Handler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	// You need to create a new ReadCloser for the request body, because it was exhausted by io.ReadAll
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var chatRequest model.CreateChatCompletionRequest
	err = c.ShouldBindBodyWith(&chatRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error binding request body"})
		return
	}
	openAiRequestUri := c.Request.RequestURI

	fmt.Println("openAiRequestUri", openAiRequestUri)

	if err != nil {
		fmt.Println(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
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
