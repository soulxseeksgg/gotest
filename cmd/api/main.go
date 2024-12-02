package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/soulxseeksgg/go-project/internal/database"
	"github.com/soulxseeksgg/go-project/internal/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("gg x")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		server.ApiConfig()
	}()

	database.DatabaseConfig()

	wg.Wait()
}
