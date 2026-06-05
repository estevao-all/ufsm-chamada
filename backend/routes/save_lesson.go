package routes

import (
	db "backend/database"
	"backend/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SaveLesson(w http.ResponseWriter, r *http.Request) {

	classId := r.PathValue("classId")

	if classId == "" {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "No id specified.")
	}

	students, err := db.FetchStudentinfo(classId)

	req_body := url.Values{}

	req_body.Set("itemDiario.dataInicio", time.Now().Format("02/01/2006")+"13:30")
	req_body.Set("itemDiario.quantidadeAulas", "2")
	req_body.Set("itemDiario.tipo", "8")
	req_body.Set("itemDiario.texto", "<p>aula</p>")
	req_body.Set("disciplina", "112944")
	req_body.Set("turmaOrigem", "983071")
	req_body.Set("origem", "D")
	req_body.Set("ano", "2026")
	req_body.Set("periodo", "101")
	req_body.Set("outrosProfessores", "true")

	for _, student := range students {
		for i := range 8 {
			req_body.Set(fmt.Sprintf("presencas[%s][%d].presente", student.Id, i), "true")
			req_body.Set(fmt.Sprintf("_presencas[%s][%d].presente", student.Id, i), "on")
		}
	}

	req_body.Set("save", "")

	req, err := http.NewRequest("POST", `https://portal.ufsm.br/docente/diario/form.html?turma=`+classId, strings.NewReader(req_body.Encode()))
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating save lesson request: "+err.Error())
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
			http.StatusInternalServerError, "Error saving lesson: "+err.Error())
		return
	}

	defer resp.Body.Close()
}
