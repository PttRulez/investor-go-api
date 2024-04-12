package router

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func getUserIdFrowJwt(r *http.Request) int {
	_, claims, _ := jwtauth.FromContext(r.Context())
	return int(claims["id"].(float64))
}
