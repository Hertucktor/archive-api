package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}
