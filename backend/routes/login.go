package routes

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var BASE_URL = "https://portal.ufsm.br"
var LOGIN_URL = BASE_URL + "/estudantil/j_security_check"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var login_request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login_request); err != nil {
		utils.WriteError(w, "invalid JSON", http.StatusBadRequest)
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

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", BASE_URL)
	req.Header.Set("Referer", LOGIN_URL)
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		error_msg := "Error making UFSM login request: " + err.Error()
		log.Println(error_msg)
		utils.WriteError(w, error_msg, http.StatusUnauthorized)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		error_msg := "Unexpected status code from UFSM login request: " + resp.Status
		log.Println(error_msg)
		utils.WriteError(w, error_msg, http.StatusUnauthorized)
		return
	}

	for _, c := range resp.Cookies() {
		log.Printf("Cookie: %s=%s\n", c.Name, c.Value)

		c.MaxAge = 60 * 60 * 24 * 7
		http.SetCookie(w, c)
	}

	w.WriteHeader(http.StatusOK)
}
