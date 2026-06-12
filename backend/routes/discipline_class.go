package routes

import (
	"backend/utils"
	"errors"
	"net/http"
	"regexp"
)

type Student struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	EnrollmentId string `json:"enrollmentId"`
}

var discipline_id_regex = regexp.MustCompile(`(?s)name="disciplina".*?value="(\d+?)"`)
var discipline_name_regex = regexp.MustCompile(`(?s)href="/docente/turma/turma\.html\?id=\d+?&action=view" class="link">\s*(.+?)\s*</a>`)
var class_name_regex = regexp.MustCompile(`(?s)href="/docente/turma/turma\.html\?id=\d+?&action=view".+?Curso.+?style="">\s*(.+?)\s*</span>`)
var default_lesson_start_time_regex = regexp.MustCompile(`(?s)name="itemDiario.dataInicio".*?value="([^"]+?)"`)
var student_regex = regexp.MustCompile(`(?s)class="link aluno" data-id="([^"]+?)">\s*(.+?)\s*</a>.+?Matrícula.+?style="">\s*(.+?)\s*</span>`)

func GetDisciplineIdFromHTML(class_form_html_content string) (string, error) {
	discipline_id_match := discipline_id_regex.FindStringSubmatch(class_form_html_content)
	if len(discipline_id_match) < 2 {
		return "", errors.New("Error parsing UFSM Portal class form HTML: Discipline ID not found in " + class_form_html_content)
	}

	return discipline_id_match[1], nil
}

func GetDisciplineNameFromHTML(class_form_html_content string) (string, error) {
	discipline_name_match := discipline_name_regex.FindStringSubmatch(class_form_html_content)
	if len(discipline_name_match) < 2 {
		return "", errors.New("Error parsing UFSM Portal class form HTML: Discipline name not found in " + class_form_html_content)
	}

	return discipline_name_match[1], nil
}

func GetClassNameFromHTML(class_form_html_content string) (string, error) {
	class_name_match := class_name_regex.FindStringSubmatch(class_form_html_content)
	if len(class_name_match) < 2 {
		return "", errors.New("Error parsing UFSM Portal class form HTML: Class name not found in " + class_form_html_content)
	}

	return class_name_match[1], nil
}

func GetDefaultLessonStartTimeFromHTML(class_form_html_content string) (string, error) {
	default_lesson_start_time_match := default_lesson_start_time_regex.FindStringSubmatch(class_form_html_content)
	if len(default_lesson_start_time_match) < 2 {
		return "", errors.New("Error parsing UFSM Portal class form HTML: Default lesson start time not found in " + class_form_html_content)
	}

	return default_lesson_start_time_match[1], nil
}

func GetStudentsFromHTML(class_form_html_content string) ([]Student, error) {
	student_matches := student_regex.FindAllStringSubmatch(class_form_html_content, -1)
	if len(student_matches) < 1 {
		return nil, errors.New("Error parsing UFSM Portal class form HTML: No students matched in " + class_form_html_content)
	}

	students := make([]Student, len(student_matches))
	for i, match := range student_matches {
		students[i] = Student{
			Id:           match[1],
			Name:         match[2],
			EnrollmentId: match[3],
		}
	}

	return students, nil
}

type DisciplineClassResponse struct {
	DisciplineId           string    `json:"disciplineId"`
	DisciplineName         string    `json:"disciplineName"`
	ClassName              string    `json:"className"`
	DefaultLessonStartTime string    `json:"defaultLessonStartTime"`
	Students               []Student `json:"students"`
}

func HandleDisciplineClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	class_id := r.PathValue("classId")
	if class_id == "" {
		utils.WriteStatusAndLogInternally(w, http.StatusBadRequest, "classId URL path parameter is required")
		return
	}

	req, err := http.NewRequest("GET", UFSM_PORTAL_CLASS_FORM_URL+class_id, nil)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating UFSM Portal class form request: "+err.Error())
		return
	}

	req.Header.Set("Referer", UFSM_PORTAL_CLASS_VIEW_URL+class_id)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making UFSM Portal class form request: "+err.Error())
		return
	}

	defer resp.Body.Close()
	if utils.HandleUnauthorized(w, r, resp) {
		return
	}

	resp_body, err := utils.ReadResponseBodyAsString(resp)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading UFSM Portal class form response body: "+err.Error())
		return
	}

	discipline_id, err := GetDisciplineIdFromHTML(resp_body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	discipline_name, err := GetDisciplineNameFromHTML(resp_body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	class_name, err := GetClassNameFromHTML(resp_body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	default_lesson_start_time, err := GetDefaultLessonStartTimeFromHTML(resp_body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	students, err := GetStudentsFromHTML(resp_body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, DisciplineClassResponse{
		DisciplineId:           discipline_id,
		DisciplineName:         discipline_name,
		ClassName:              class_name,
		DefaultLessonStartTime: default_lesson_start_time,
		Students:               students,
	})
}
