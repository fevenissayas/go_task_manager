package infrastructure

import (
	"net/http"
	"restfulapi/Usecases"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSecret = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["user_id"].(string); ok {
				ctx.Set("userID", userID)
			}
		}

		ctx.Next()
	}
}

func AdminOnly(userUsecase *usecases.UserUsecase) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userIDstr, exists := ctx.Get("userID")
        if !exists {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
            ctx.Abort()
            return
        }
        objectID, err := primitive.ObjectIDFromHex(userIDstr.(string))
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
            ctx.Abort()
            return
        }

        user, err := userUsecase.GetUserByID(objectID)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            ctx.Abort()
            return
        }
        if user.Role != "admin" {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}