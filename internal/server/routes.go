package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func ApiConfig() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", port),
		Handler:      r,
		IdleTimeout:  5 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
