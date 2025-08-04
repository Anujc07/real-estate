package jwtutil

import (
    "time"
    "os"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, role string) (string, error) {
    claims := jwt.MapClaims{
        "id":   userID,
        "role": role,
        "exp":  time.Now().Add(time.Hour * 24).Unix(), // expires in 24hr
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
