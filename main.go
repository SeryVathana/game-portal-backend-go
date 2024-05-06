package main

import (
	"fmt"

	"direxplay.com/game-portal-backend-go/core"
	"github.com/gin-gonic/gin"
)

func main() {
	config := core.LoadConfig()
	// Connect to database
	core.ConnectToDB()

	router := gin.Default()
	router.SetTrustedProxies([]string{})

	fmt.Println("Server is running...")

	router.Run(fmt.Sprintf(":%d", config.Port))
}
