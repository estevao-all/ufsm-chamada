package routes

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var login_request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login_request); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid LoginRequest body")
		return
	}

	req_body := url.Values{}
	req_body.Set("j_username", login_request.Username)
	req_body.Set("j_password", login_request.Password)
	req_body.Set("enter", "")

	req, err := http.NewRequest("POST", UFSM_PORTAL_LOGIN_URL, strings.NewReader(req_body.Encode()))
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating UFSM Portal login request: "+err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Origin", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Referer", UFSM_PORTAL_INDEX_URL)
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making UFSM Portal login request: "+err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized) // Invalid credentials.
		return
	}

	if resp.StatusCode != http.StatusFound {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Unexpected status code from UFSM Portal login request: "+resp.Status)
		return
	}

	utils.RedirectCookiesAndSetMaxAge(w, resp)
	for _, cookie := range resp.Cookies() {
		log.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	}

	w.WriteHeader(http.StatusOK) // Login successful.
}
