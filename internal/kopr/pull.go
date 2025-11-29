package kopr

import (
	"encoding/json"
	"log"
	"net/http"
)

type authToken struct {
	Token string `json:"token"`
}

func GetAuthToken(baseURL string, repo string) string {
	url := baseURL + "/token?service=registry.docker.io&scope=repository:" + repo + ":pull"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var auth authToken
	json.NewDecoder(response.Body).Decode(&auth)

	return auth.Token
}
