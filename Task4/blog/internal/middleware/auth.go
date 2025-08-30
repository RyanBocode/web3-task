package middleware

import (
	"blog/internal/config"
	"blog/internal/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	CTXUserIDKey   = "userID"
	CTXUsernameKey = "username"
)

func NewToken(cfg *config.Config, uid uint, username string) (string, error) {
	claims := &Claims{
		UserID:   uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWTTTLHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatUint(uint64(uid), 10),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.JWTSecret))
}

func AuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ah := c.GetHeader("Authorization")
		if len(ah) < 8 || ah[:7] != "Bearer " {
			responses.JSONError(c, http.StatusUnauthorized, "unauthorized", "missing or invalid Authorization header")
			return
		}
		tokenStr := ah[7:]
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			responses.JSONError(c, http.StatusUnauthorized, "unauthorized", "invalid token")
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok {
			responses.JSONError(c, http.StatusUnauthorized, "unauthorized", "invalid claims")
			return
		}
		c.Set(CTXUserIDKey, claims.UserID)
		c.Set(CTXUsernameKey, claims.Username)
		c.Next()
	}
}

func MustGetUserID(c *gin.Context) uint {
	if v, exists := c.Get(CTXUserIDKey); exists {
		if id, ok := v.(uint); ok {
			return id
		}
	}
	return 0
}
