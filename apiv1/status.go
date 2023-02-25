package apiv1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func statusAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func statusCheck(c *gin.Context) {
	responseWriter := c.Writer
	c.Header("Content-Type", "application/json")
	body := struct {
		ResponseCode int
	}{
		ResponseCode: http.StatusOK,
	}

	marshalledObject, err := json.Marshal(body)
	if err != nil {
		log.Printf("Couldn't marshal body %v", err)

		responseWriter.WriteHeader(http.StatusInternalServerError)
		if _, err = responseWriter.Write([]byte("something bad happened, please contact the administrator")); err != nil {
			log.Fatalf("Couldn't write error message back to user %v", err)
		}
		return
	}

	responseWriter.WriteHeader(body.ResponseCode)
	if _, err = responseWriter.Write(marshalledObject); err != nil {
		log.Fatalf("Couldn't write info response back to user %v", err)
	}
}
