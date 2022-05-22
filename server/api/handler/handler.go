package handler

import (
	apimiddleware "advdiploma/server/api/middleware"
	"advdiploma/server/api/model"
	"advdiploma/server/services/auth"
	"advdiploma/server/services/secret"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

type Handler struct {
	svcAuth   auth.Authenticator
	svcSecret secret.SecretManager
}

// NewHandler Return new handler
func NewHandler(auth auth.Authenticator, secret secret.SecretManager) (*Handler, error) {

	return &Handler{
		svcAuth:   auth,
		svcSecret: secret,
	}, nil
}

func GetRouter(handler *Handler) *chi.Mux {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	r := chi.NewRouter()
	r.Use(middleware.Compress(5))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		//r.Post("/api/user/register", handler.Register)
		//r.Post("/api/user/login", handler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(apimiddleware.MiddlewareAuth)

		// Secret processing
		r.Put("/api/secret", handler.SecretAdd)
		//r.Get("/api/secret", handler.OrderGetListForUser)
		//r.Delete("/api/secret", handler.OrderAddToUser)

	})

	return r
}

func (h *Handler) GetUserDataFromContext(r *http.Request) model.UserContextData {
	ctxData := r.Context().Value(model.ContextKeyUserID).(model.UserContextData)

	return ctxData
}
