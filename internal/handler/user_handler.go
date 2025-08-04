package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

    "github.com/Anujc07/real-estate/config"
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
    "github.com/Anujc07/real-estate/pkg/jwt"
)

func RegisterUserRoutes(r *gin.Engine) {
    userRepo := repository.NewUserRepository(config.DB)

    r.POST("/register", func(c *gin.Context) {
        var input models.User
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
        input.Password = string(hashed)

        if err := userRepo.Create(&input); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
    })

    r.POST("/login", func(c *gin.Context) {
        var input models.User
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        user, err := userRepo.FindByEmail(input.Email)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        token, err := jwtutil.GenerateJWT(user.ID, user.Role)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    })
}
