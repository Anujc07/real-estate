package handler

import (
    "strconv"
    "strings"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

    "github.com/Anujc07/real-estate/config"
    "github.com/Anujc07/real-estate/internal/models"
    "github.com/Anujc07/real-estate/internal/repository"
    "github.com/Anujc07/real-estate/pkg/jwt"
)

// RegisterUserRoutes sets up user-related routes for registration and login
func RegisterUserRoutes(r *gin.Engine) {
    // Initialize the user repository with the database connection
    userRepo := repository.NewUserRepository(config.DB)

    // ------------------ User Registration Route ------------------
    r.POST("/register", func(c *gin.Context) {
        var input models.User

        // Parse and bind incoming JSON request to User model
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Set default role if not provided
        if input.Role == "" {
            input.Role = "customer"
        }

        // Validate role against allowed roles (admin, customer, seller)
        if !models.AllowedRoles[input.Role] {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Allowed roles: admin, customer, seller"})
            return  
        }

        // Hash the password before storing
        hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
        input.Password = string(hashed)

        // Save the new user to the database
        if err := userRepo.Create(&input); err != nil {

            // ✅ Check if the error message mentions duplicate key
            if strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "email") {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
            return
        }

        // Success response
        c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
    })

    // ------------------ User Login Route ------------------
    r.POST("/login", func(c *gin.Context) {
        var input models.User

        // Parse and bind incoming JSON request to User model
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Fetch user by email from the database
        user, err := userRepo.FindByEmail(input.Email)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if !user.IsActive {
            c.JSON(http.StatusForbidden, gin.H{"error": "User is inactive. Please contact support."})
            return
        }

        // Compare provided password with stored hashed password
        if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        // Generate JWT token with user ID and role
        token, err := jwtutil.GenerateJWT(user.ID, user.Role)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        // Success response with token
        c.JSON(http.StatusOK, gin.H{"token": token})
    })

    r.POST("/user/in-active/:id", func(c *gin.Context) {
        // Get ID from path param
        idParam := c.Param("id")

        // Convert string to uint
        id, err := strconv.ParseUint(idParam, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }

        // Call repository to deactivate user
        err = userRepo.DeactivateUserByID(uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate user"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User deactivated successfully"})
    })

    r.POST("/user/active/:id", func(c *gin.Context) {
        // Get ID from path param
        idParam := c.Param("id")

        // Convert string to uint
        id, err := strconv.ParseUint(idParam, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }

        // Call repository to activate user
        err = userRepo.ActivateUserByID(uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to activate user"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User activated successfully"})
    })

}
