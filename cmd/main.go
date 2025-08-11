package main

import (
    "log"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/Anujc07/real-estate/config"
    "github.com/Anujc07/real-estate/internal/handler"
    "github.com/Anujc07/real-estate/internal/routes"
)

func main() {
    config.InitDB()

    r := gin.Default()

    // ✅ Enable CORS for frontend on port 3000
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    handler.RegisterUserRoutes(r)
    routes.SetupRoutes(r)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
