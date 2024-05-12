package middlewares

import (
	"net/http"
	"tanam-backend/helpers"

	"github.com/labstack/echo/v4"
)

func ExtractTokenFromHeader(authHeader string) string {
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}
	return ""
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := ExtractTokenFromHeader(c.Request().Header.Get("Authorization"))
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed token")
		}

		// Validate token
		user, err := helpers.ValidateToken(token, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		c.Set("user", user)

		return next(c)
	}
}
