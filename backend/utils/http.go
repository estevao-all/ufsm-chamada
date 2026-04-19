package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var default_cookies_max_age = 60 * 60 * 24 * 7

var Client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ReadResponseBodyAsString(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func WriteStatusAndLogInternally(w http.ResponseWriter, statusCode int, message string) {
	log.Println(message)
	w.WriteHeader(statusCode)
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, ErrorResponse{Error: message})
}

func WriteErrorAndLogInternally(w http.ResponseWriter, statusCode int, message string) {
	log.Println(message)
	WriteError(w, statusCode, message)
}

func RedirectCookiesAndSetMaxAge(w http.ResponseWriter, resp *http.Response) {
	for _, cookie := range resp.Cookies() {
		cookie.Path = "/"
		cookie.HttpOnly = false
		cookie.MaxAge = default_cookies_max_age
		http.SetCookie(w, cookie)
	}
}

func ClearCookies(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		http.SetCookie(w, &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Path:     "/",
			HttpOnly: false,
			MaxAge:   -1,
		})
	}
}

func Unauthorize(w http.ResponseWriter, r *http.Request) {
	ClearCookies(w, r)
	WriteError(w, http.StatusUnauthorized, "Unauthorized")
}

func HandleUnauthorized(w http.ResponseWriter, r *http.Request, resp *http.Response) bool {
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		Unauthorize(w, r)
		return true
	}

	return false
}
