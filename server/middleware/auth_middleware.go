package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
)

func trimPrefix(token string) string {
	prefix := "Bearer "
	if strings.HasPrefix(token, prefix) {
		return strings.TrimPrefix(token, prefix)
	}

	return ""
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			err := rest.NewUnauthorizedErr()
			rest.JSON(w, err.Code, err)
			return
		}

		token, err := jwt.Parse(trimPrefix(authorization), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(config.GetConfig().JWT), nil
			}

			return nil, rest.NewUnauthorizedErr()
		})

		if err != nil {
			rest_err := rest.NewUnauthorizedErr()
			rest.JSON(w, rest_err.Code, rest_err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			err := rest.NewUnauthorizedErr()
			rest.JSON(w, err.Code, err)
			return
		}

		r.Header.Set("X-Authenticated-Employee", claims["registry"].(string))

		next.ServeHTTP(w, r)
	})
}
