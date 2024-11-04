package server

import (
	"context"
	"fmt"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"net/http"
)

func Run(ctx context.Context) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithOAuth("google",
				config.WithAllowOAuthEmail("tranvuquanganh87@gmail.com"),
			),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("App URL", listener.URL())
	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello from ngrok-go!</h1>")

	r.Header.Values("ngrok-auth-oauth-email")
}
