package routes

import (
	"backend/database"
	db "backend/database"
	"backend/models"
	"backend/utils"
	"database/sql"
	"errors"
	"io"
	"net/http"
	"regexp"
)

var DB *sql.DB

type StudentInfo = models.StudentInfo

type DisciplineStudentsResponse struct {
	Students []StudentInfo `json:"students"`
}

var studentPattern = regexp.MustCompile(`(?s)class="link aluno" data-id="([^"]+?)">([^<]+).*?Matrícula.*?style="">(.+?)<`)

func ParseStudents(htmlContent string) ([]StudentInfo, error) {
	matches := studentPattern.FindAllStringSubmatch(htmlContent, -1)

	if len(matches) < 2 {
		return nil, errors.New("Error parsing user info HTML: Error matching pattern")
	}

	students := make([]StudentInfo, len(matches))

	for i, match := range matches {
		students[i] = StudentInfo{
			Id:        match[1],
			Name:      match[2],
			Matricula: match[3],
		}
	}

	return students, nil
}

func CreateLesson(w http.ResponseWriter, r *http.Request) {

	classId := r.PathValue("classId")

	if classId == "" {
		utils.WriteStatusAndLogInternally(w, http.StatusBadRequest, "classId parameter is required")
		return
	}

	req, err := http.NewRequest("GET", `https://portal.ufsm.br/docente/diario/form.html?turma=`+classId, nil)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating lesson request: "+err.Error())
		return
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Origin", UFSM_PORTAL_BASE_URL)
	req.Header.Set("Referer", UFSM_CREATE_LESSOR_REFERER+classId)
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making Discipline Student request: "+err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading Discipline Student request "+err.Error())
	}

	students, err := ParseStudents(string(body))
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	db.WriteStudentInfo(students, classId)

	utils.WriteJSON(w, http.StatusOK, DisciplineStudentsResponse{Students: students})
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	classId := r.PathValue("classId")
	s, _ := database.FetchStudentinfo(classId)

	utils.WriteJSON(w, http.StatusOK, DisciplineStudentsResponse{Students: s})

}
