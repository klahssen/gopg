package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	}
	mid := authHeaderMiddleware(http.HandlerFunc(h))
	tests := []struct {
		authHeader string
		expCode    int
		//expBody string
	}{
		{
			authHeader: "", expCode: http.StatusUnauthorized,
		},
		{
			authHeader: "Bearer Kl09i009312lkdasja9-kdjskdj", expCode: http.StatusOK,
		},
	}

	for ind, test := range tests {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "http://localhost:8080/protected", nil)
		if test.authHeader != "" {
			r.Header.Set("Authorization", "Bearer Kl09i009312lkdasja9-kdjskdj")
		}
		mid.ServeHTTP(w, r)
		if w.Code != test.expCode {
			t.Errorf("test %d: expected code %d received %d", ind, test.expCode, w.Code)
		}
	}

}
