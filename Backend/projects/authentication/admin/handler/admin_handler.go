package handler

import (
	"auth_admin/domain"
	"auth_admin/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminUsecase domain.AdminUsecase
}

func NewAdminHandler(e *echo.Echo, as domain.AdminUsecase) {
	handler := &AdminHandler{
		AdminUsecase: as,
	}

	e.POST("/admin/login", handler.Login)
	e.POST("/admin/logout", handler.Logout, middleware.JWTMiddleware)

}

func (h *AdminHandler) Login(c echo.Context) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := new(request)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.AdminUsecase.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	type response struct {
		Token string `json:"token"`
		Data  struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"data"`
	}

	res := response{
		Token: token,
	}

	res.Data.Username = req.Username
	res.Data.Password = req.Password

	return c.JSON(http.StatusOK, res)
}

func (h *AdminHandler) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully logged out",
	})
}
