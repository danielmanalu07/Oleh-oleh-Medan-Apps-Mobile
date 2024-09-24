package handler

import (
	"auth_admin/infrastucture/middleware"
	"auth_admin/service"
	"fmt"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	tokenBlacklist = make(map[string]struct{})
	mu             sync.Mutex
)

type AdminHandler struct {
	AdminService *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{AdminService: service}
}

func (h *AdminHandler) Login(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "username and password are required",
		})
	}

	admin, err := h.AdminService.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	token, err := middleware.GenerateToken(admin.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Could not generate token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"data": map[string]string{
			"id":       fmt.Sprintf("%d", admin.Id),
			"username": admin.Username,
		},
	})
}

func (h *AdminHandler) Logout(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized",
		})
	}

	// Type assertion untuk mendapatkan token JWT
	token, ok := user.(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid token",
		})
	}

	// Ambil token string mentah
	tokenString := token.Raw

	// Tambahkan token ke blacklist
	mu.Lock()
	tokenBlacklist[tokenString] = struct{}{}
	mu.Unlock()

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logged out",
	})
}
