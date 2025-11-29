package kopr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetToken(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scope := r.PathValue("scope")

		if scope != "repository:ubuntu:pull" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"token\":\"XASDXSADSADSXASXASDASDASDAXXSADASDASXASDASDASXASDASDADSXASDASDASXASDASDASXXSXXASDSAXAXSXF\"}"))
	}))
	defer testServer.Close()

	token := GetAuthToken(testServer.URL, "ubuntu")

	if len(token) < 50 {
		t.Fatal("Something wrong with retrieved token")
	}
}
