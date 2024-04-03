package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", responseToAddress)

	r.Run(":8080")
}

func responseToAddress(ctx *gin.Context) {
	type RequestBody struct {
		From string `json:from`
	}

	var requestBody RequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	from := requestBody.From

	fmt.Println("From: ", from)

	time.Sleep(1 * time.Second)

	responseData := map[string]bool{
		"web2response": true,
	}

	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

	ctx.Data(http.StatusOK, "application/json", jsonResponse)
}