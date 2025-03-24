package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/movie-app-crud-gorm/internal/pkg/status"

	"github.com/movie-app-crud-gorm/internal/domain"
	"net/http"
	"os"
	"strings"
)

type TokenData struct {
	UserId string `json:"user_id"`
}

type JwtMiddleware struct {
	userRepo domain.UserRepository
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{}
}

func (m *JwtMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk, err := extractTokenMetaData(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, R{
				Status:    status.Failure,
				ErrorCode: status.ErrorCodeValidation,
				ErrorNote: "Not Authorized",
			})
			c.Abort()
			return
		}

		user, errU := m.userRepo.GetByID(c.Request.Context(), tk.UserId)
		if errU != nil {
			c.JSON(http.StatusUnauthorized, errView(status.ErrorUnauthorizedAccess, errU.Error()))
			c.Abort()
			return
		}

		c.Set("user_id", user.ID)
		c.Next()
	}
}

func extractTokenMetaData(r *http.Request) (*TokenData, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, _ := claims["user_id"].(string)

		return &TokenData{
			UserId: userID,
		}, nil
	}
	return nil, errors.New("invalid token claims")
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	if tokenString == "" {
		return nil, errors.New("no token provided")
	}
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func extractToken(r *http.Request) string {
	b := r.Header.Get("Authorization")
	parts := strings.Split(b, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
