package routes

import (
    "github.com/gin-gonic/gin"
    "user-service/controllers"
)

func UserRoutes(r *gin.Engine) {
    user := r.Group("/users")
    {
        user.POST("/", controllers.CreateUser)
        user.GET("/", controllers.GetUsers)
    }
}
