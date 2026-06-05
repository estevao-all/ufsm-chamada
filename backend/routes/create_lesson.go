package routes

import (
	"backend/utils"
	"net/http"
)

func CreateLesson(w http.ResponseWriter, r *http.Request) {

	classId := r.PathValue("classId")

	if classId == "" {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "No id specified.")
	}

	req, err := http.NewRequest("GET", `https://portal.ufsm.br/docente/diario/form.html?turma=`+classId, nil)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating lesson request: "+err.Error())
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Origin", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Referer", UFSM_CREATE_LESSOR_REFERER+classId)
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making lesson creation request: "+err.Error())
	}

	defer resp.Body.Close()

	resp_body, err := utils.ReadResponseBodyAsString(resp)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading created lesson body: "+err.Error())
	}

	// parse body

	// write to database

	utils.WriteJSON(w, http.StatusOK, resp_body)

}
