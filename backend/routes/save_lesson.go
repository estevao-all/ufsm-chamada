package routes

import (
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type StudentPresence struct {
	StudentId string `json:"studentId"`
	Status    bool   `json:"status"`
}

type SaveLessonRequest struct {
	DisciplineId     string            `json:"disciplineId"`
	StartTime        string            `json:"startTime"`
	HourAmount       string            `json:"hourAmount"`
	Type             string            `json:"type"`
	NoteText         string            `json:"noteText"`
	RemoteLesson     bool              `json:"remoteLesson"`
	Coil             bool              `json:"coil"`
	StudentPresences []StudentPresence `json:"studentPresences"`
}

func HandleSaveLesson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	class_id := r.PathValue("classId")
	if class_id == "" {
		utils.WriteStatusAndLogInternally(w, http.StatusBadRequest, "classId URL path parameter is required")
		return
	}

	var save_lesson_request SaveLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&save_lesson_request); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid SaveLessonRequest body")
		return
	}

	req_body := url.Values{}
	req_body.Set("itemDiario.dataInicio", save_lesson_request.StartTime)
	req_body.Set("itemDiario.quantidadeAulas", save_lesson_request.HourAmount)
	req_body.Set("itemDiario.tipo", save_lesson_request.Type)
	req_body.Set("itemDiario.texto", "<p>"+save_lesson_request.NoteText+"</p>")
	req_body.Set("itemDiario.ead", strconv.FormatBool(save_lesson_request.RemoteLesson))
	req_body.Set("itemDiario.coil", strconv.FormatBool(save_lesson_request.Coil))
	req_body.Set("disciplina", save_lesson_request.DisciplineId)
	req_body.Set("turmaOrigem", class_id)
	req_body.Set("origem", "D")
	req_body.Set("ano", "2026")
	req_body.Set("periodo", "101")
	req_body.Set("outrosProfessores", "true")

	for _, student_presence := range save_lesson_request.StudentPresences {
		for i := range 8 { // 8 is the maximum amount of hours a lesson can have.
			if student_presence.Status {
				req_body.Set(fmt.Sprintf("presencas[%s][%d].presente", student_presence.StudentId, i), "true")
			}

			req_body.Set(fmt.Sprintf("_presencas[%s][%d].presente", student_presence.StudentId, i), "on")
		}
	}

	req_body.Set("save", "")

	req, err := http.NewRequest("POST", UFSM_PORTAL_CLASS_FORM_URL+class_id, strings.NewReader(req_body.Encode()))
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating UFSM Portal save lesson request: "+err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Referer", UFSM_PORTAL_CLASS_FORM_URL+class_id)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making UFSM Portal save lesson request: "+err.Error())
		return
	}

	if resp.StatusCode != http.StatusFound {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Unexpected status code from UFSM Portal save lesson request: "+resp.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
}
