package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
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
	reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func WriteStatusAndLogInternally(w http.ResponseWriter, status_code int, message string) {
	log.Println(message)
	w.WriteHeader(status_code)
}

func WriteJSON(w http.ResponseWriter, status_code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status_code int, message string) {
	WriteJSON(w, status_code, ErrorResponse{Error: message})
}

func WriteErrorAndLogInternally(w http.ResponseWriter, status_code int, message string) {
	log.Println(message)
	WriteError(w, status_code, message)
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
