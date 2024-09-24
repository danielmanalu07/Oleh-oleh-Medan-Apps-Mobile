package middleware

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	tokenBlacklist = make(map[string]struct{})
	mu             sync.Mutex
)

func InitJWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret_key"),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ErrorHandler: func(err error) error {
			return echo.ErrUnauthorized
		},
	})
}

func CheckBlacklist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		if user == nil {
			return echo.ErrUnauthorized
		}

		token, ok := user.(*jwt.Token)
		if !ok {
			return echo.ErrUnauthorized
		}

		tokenString := token.Raw

		mu.Lock()
		_, found := tokenBlacklist[tokenString]
		mu.Unlock()

		if found {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func GenerateToken(adminID uint) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret_key"))
}
