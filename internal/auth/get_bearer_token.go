package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	authList := strings.Split(auth, " ")

	if len(authList) != 2 || authList[1] == "" || authList[0] != "Bearer" {
		return "", fmt.Errorf("couldn't find authorization header.")
	}
	return authList[2], nil
}
