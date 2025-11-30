package kopr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func validateURLParam(key string, expectedValue string, request *http.Request, writer http.ResponseWriter) {
	actualValue := request.URL.Query().Get(key)

	if actualValue != expectedValue {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func TestGetToken(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validateURLParam("scope", "repository:ubuntu:pull", r, w)
		validateURLParam("service", "registry.docker.io", r, w)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"token\":\"XASDXSADSADSXASXASDASDASDAXXSADASDASXASDASDASXASDASDADSXASDASDASXASDASDASXXSXXASDSAXAXSXF\"}"))
	}))
	defer testServer.Close()

	token := GetAuthToken(testServer.URL, "ubuntu")

	if len(token) < 50 {
		t.Fatal("Something wrong with retrieved token")
	}
}
