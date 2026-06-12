package routes

import (
	"backend/utils"
	"io"
	"net/http"
	"strings"
)

func HandleTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req_body := strings.Join([]string{
		"callCount=1",
		"nextReverseAjaxIndex=0",
		"c0-scriptName=gradeHorariosAjaxService",
		"c0-methodName=horarios",
		"c0-id=0",
		"c0-param0=string:2026",
		"c0-param1=string:101",
		"batchId=0",
		"instanceId=0",
		"page=%2Fdocente%2Fturma%2Fturma.html%3Faction%3Dlist",
		"scriptSessionId=YTYvGc~Ngqt81Iu0C53vM!YU!lHzrdsIgVp/MysIgVp-EK2*u7G*8",
	}, "\n")

	req, err := http.NewRequest("POST", UFSM_PORTAL_SCHEDULE_URL, strings.NewReader(req_body))
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating UFSM Portal teacher schedule request: "+err.Error())
		return
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Origin", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Referer", UFSM_PORTAL_CLASSES_URL)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making UFSM Portal teacher schedule request: "+err.Error())
		return
	}

	defer resp.Body.Close()
	if utils.HandleUnauthorized(w, r, resp) {
		return
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error copying UFSM Portal teacher schedule response body "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/javascript;charset=ISO-8859-1")
}
