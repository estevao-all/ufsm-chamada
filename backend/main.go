package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/login", handleLogin)

	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "./frontend" + r.URL.Path

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, "./frontend/index.html")
			return
		}

		fs.ServeHTTP(w, r)
	}))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, message string, status int) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var login_request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login_request); err != nil {
		writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	ufsm_login_body := url.Values{}
	ufsm_login_body.Set("j_username", login_request.Username)
	ufsm_login_body.Set("j_password", login_request.Password)
	ufsm_login_body.Set("enter", "")

	req, _ := http.NewRequest(
		"POST",
		"https://portal.ufsm.br/estudantil/j_security_check",
		strings.NewReader(ufsm_login_body.Encode()),
	)

	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://portal.ufsm.br")
	req.Header.Set("Referer", "https://portal.ufsm.br/estudantil/j_security_check")

	resp, err := client.Do(req)
	if err != nil {
		error_msg := "Error making UFSM login request: " + err.Error()
		log.Println(error_msg)
		writeError(w, error_msg, http.StatusUnauthorized)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		error_msg := "Unexpected status code from UFSM login request: " + resp.Status
		log.Println(error_msg)
		writeError(w, error_msg, http.StatusUnauthorized)
		return
	}

	for _, c := range resp.Cookies() {
		log.Printf("Cookie: %s=%s\n", c.Name, c.Value)

		c.MaxAge = 60 * 60 * 24 * 7
		http.SetCookie(w, c)
	}

	w.WriteHeader(http.StatusOK)
}
