package main

import (
    "github.com/gin-gonic/gin"
    "user-service/config"
    "user-service/routes"
)

func main() {
    r := gin.Default()

    config.ConnectDatabase()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "User Service is running!"})
    })

    routes.UserRoutes(r)

    r.Run(":8001") // Listen on port 8001
}
