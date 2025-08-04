package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/Anujc07/real-estate/config"
    "github.com/Anujc07/real-estate/internal/handler"
    "github.com/Anujc07/real-estate/internal/routes"
)

func main() {
    config.InitDB()

    r := gin.Default()

    handler.RegisterUserRoutes(r)
    routes.SetupRoutes(r)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
