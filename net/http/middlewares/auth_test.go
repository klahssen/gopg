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
	/*srv := httptest.NewServer(authHeaderMiddleware(h))
		defer srv.Close()

		tests:=[]struct{
	url string
		}{
			{}
		}
		http.Get(test.url)
	*/

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	r1 := httptest.NewRequest("GET", "http://localhost:8080/protected", nil)
	r2 := httptest.NewRequest("GET", "http://localhost:8080/protected", nil)
	r2.Header.Set("Authorization", "Bearer Kl09i009312lkdasja9-kdjskdj")
	mid := authHeaderMiddleware(http.HandlerFunc(h))
	mid.ServeHTTP(w1, r1)
	if w1.Code != http.StatusUnauthorized {
		t.Errorf("request without Authorization header did not fail!")
	}
	mid.ServeHTTP(w2, r2)
	if w2.Code != http.StatusOK {
		t.Errorf("request with Authorization header is not OK!")
	}
}
