package routes

import (
	"backend/utils"
	"net/http"
	"regexp"
)

type UserInfoResponse struct {
	Name string `json:"name"`
}

var name_regex = regexp.MustCompile(`Menu Geral do Usuário.+?icon-user"></span> (.+?) <span`)

func HandleUserInfo(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", UFSM_PORTAL_USER_INFO_URL, nil)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating UFSM Portal user info request: "+err.Error())
		return
	}

	req.Header.Set("Host", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making UFSM Portal user info request: "+err.Error())
		return
	}

	defer resp.Body.Close()

	if utils.HandleUnauthorized(w, r, resp) {
		return
	}

	resp_body, err := utils.ReadResponseBodyAsString(resp)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading UFSM Portal user info response body: "+err.Error())
		return
	}

	name_match := name_regex.FindStringSubmatch(resp_body)
	if len(name_match) < 2 {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError,
			"Error parsing UFSM Portal user info HTML: Name not found in"+resp_body)
		return
	}

	utils.WriteJSON(w, http.StatusOK, UserInfoResponse{Name: name_match[1]})
}
