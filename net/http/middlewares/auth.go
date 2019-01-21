package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type ctxKey string

func authHeaderMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := getAuthHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %s\n", err.Error())
			return
		}
		ctx := context.Background()
		ctx = context.WithValue(ctx, ctxKey("auth-token"), token)
		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing 'Authorization' header")
	}
	if !strings.HasPrefix(authHeader, "Bearer") {
		return "", fmt.Errorf("must satisfy Authorization Bearer scheme")
	}
	token := strings.TrimSpace(strings.TrimLeft(authHeader, "Bearer "))
	return token, nil
}
