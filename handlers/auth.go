package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	configs map[string]AuthConfig
}

type AuthConfig struct {
	OidcProvider *oidc.Provider
	OidcConfig   *oidc.Config
	Oauth2Config *oauth2.Config
}

func NewAuthHandler(authConfigs map[string]AuthConfig) *AuthHandler {
	return &AuthHandler{
		configs: authConfigs,
	}
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func (a *AuthHandler) Register(r chi.Router) {
	r.Route("/auth", func(signinRouter chi.Router) {
		for key, val := range a.configs {
			signinRouter.Route(fmt.Sprintf("/%s", key), func(serviceRouter chi.Router) {
				serviceRouter.Get("/callback", func(w http.ResponseWriter, r *http.Request) {
					state, err := r.Cookie("state")
					if err != nil {
						http.Error(w, "state not found", http.StatusBadRequest)
						return
					}
					if r.URL.Query().Get("state") != state.Value {
						http.Error(w, "state did not match", http.StatusBadRequest)
						return
					}

					oauth2Token, err := val.Oauth2Config.Exchange(r.Context(), r.URL.Query().Get("code"))
					if err != nil {
						http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
						return
					}
					rawIDToken, ok := oauth2Token.Extra("id_token").(string)
					if !ok {
						http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
						return
					}
					idToken, err := val.OidcProvider.Verifier(val.OidcConfig).Verify(r.Context(), rawIDToken)
					if err != nil {
						http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
						return
					}

					nonce, err := r.Cookie("nonce")
					if err != nil {
						http.Error(w, "nonce not found", http.StatusBadRequest)
						return
					}
					if idToken.Nonce != nonce.Value {
						http.Error(w, "nonce did not match", http.StatusBadRequest)
						return
					}
				})
				serviceRouter.Get("/signin", func(w http.ResponseWriter, r *http.Request) {
					state := ""
					nonce := ""
					setCallbackCookie(w, r, "state", state)
					setCallbackCookie(w, r, "nonce", nonce)
					http.Redirect(w, r, val.Oauth2Config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
				})
			})
		}
	})
}
