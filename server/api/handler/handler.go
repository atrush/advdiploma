package handler

import (
	apimiddleware "advdiploma/server/api/middleware"
	"advdiploma/server/services/auth"
	"advdiploma/server/services/secret"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	svcAuth   auth.Authenticator
	svcSecret secret.SecretManager
	jwtAuth   *jwtauth.JWTAuth
}

// NewHandler Return new handler
func NewHandler(auth auth.Authenticator, secret secret.SecretManager, jwtAuth *jwtauth.JWTAuth) (*Handler, error) {

	return &Handler{
		svcAuth:   auth,
		svcSecret: secret,
		jwtAuth:   jwtAuth,
	}, nil
}

func GetRouter(handler *Handler) *chi.Mux {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	r := chi.NewRouter()
	r.Use(middleware.Compress(5))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		r.Post("/api/user/register", handler.Register)
		r.Post("/api/user/login", handler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(apimiddleware.MiddlewareAuth)

		r.Put("/api/sync", handler.SyncList)

		// Secret processing
		r.Put("/api/secret", handler.SecretUpload)
		r.Get("/api/secret", handler.SecretGet)
		r.Delete("/api/secret", handler.SecretDelete)

	})

	return r
}
