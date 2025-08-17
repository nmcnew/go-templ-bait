package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nmcnew/tournament-winner/handlers"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)

func getDiscordAuthConfig(ctx context.Context) handlers.AuthConfig {
	discordClientId := os.Getenv("DISCORD_CLIENT_ID")
	discordClientSecret := os.Getenv("DICSCORD_CLIENT_SECRET")
	provider, err := oidc.NewProvider(ctx, "https://discord.com")
	if err != nil {
		slog.Error("Setting up discord oidc provider failed", "error", err)
		panic(err)
	}
	discordOidcConfig := &oidc.Config{
		ClientID: discordClientId,
	}
	discordOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/discord/signin",
		ClientID:     discordClientId,
		ClientSecret: discordClientSecret,
		Scopes:       []string{oidc.ScopeOpenID, "email", discord.ScopeIdentify},
		Endpoint:     provider.Endpoint(),
	}
	return handlers.AuthConfig{
		OidcProvider: provider,
		OidcConfig:   discordOidcConfig,
		Oauth2Config: discordOauthConfig,
	}
}
func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	handlers.NewAuthHandler(map[string]handlers.AuthConfig{
		"discord": getDiscordAuthConfig(ctx),
	})
	indexHandler := &handlers.IndexHandler{}
	indexHandler.Register(r)

	userHandler := &handlers.UserHandler{}
	userHandler.Register(r)

	http.ListenAndServe(":3000", r)
}
