package routes

import (
	"backend/utils"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type TeacherDiscipline struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Time string `json:"time"`
}

type TeacherDisciplinesResponse struct {
	Disciplines []TeacherDiscipline `json:"disciplines"`
}

var sessionIdRegex = regexp.MustCompile(`(?s)handleCallback\(\\"0\\",\\"0\\",\\"(.+?)\\"\)`)

func ParseSessionId(responseBody []byte) string {
	sessionId := sessionIdRegex.FindStringSubmatch(string(responseBody))
	if len(sessionId) < 2 {
		return ""
	}
	return sessionId[1]
}

func GetScriptSessionId(w http.ResponseWriter, r *http.Request) string {

	// "text": "callCount=1\nc0-scriptName=__System\nc0-methodName=generateId\nc0-id=0\nbatchId=0\ninstanceId=1\npage=%2Fdocente%2FmainMenu.html\nscriptSessionId=\n",

	postBody := strings.NewReader("callCount=1\n" +
		"c0-scriptName=__System\n" +
		"c0-methodName=generateId\n" +
		"c0-id=0\n" +
		"batchId=0\n" +
		"instanceId=1\n" +
		"page=%2Fdocente%2FmainMenu.html\n" +
		"scriptSessionId=\n")

	req, err := http.NewRequest("POST", "https://portal.ufsm.br/docente/dwr/call/plaincall/__System.generateId.dwr", postBody)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating Discipline request: "+err.Error())
		return ""
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Content-Type", "text/plain")

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making Discipline request: "+err.Error())
		return ""
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading Discipline request "+err.Error())
	}

	sessionId := ParseSessionId(responseBody)

	if sessionId == "" {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError,
			"Error parsing SessionId: SessionId not found in "+string(responseBody))
		return ""
	}

	return sessionId

}

func GetDisciplines(w http.ResponseWriter, r *http.Request, scriptSessionId string) {

	bodyText := "callCount=1\n" +
		"nextReverseAjaxIndex=0\n" +
		"c0-scriptName=gradeHorariosAjaxService\n" +
		"c0-methodName=horarios\n" +
		"c0-id=0\n" +
		"c0-param0=string:2026\n" +
		"c0-param1=string:101\n" +
		"batchId=0\n" +
		"instanceId=0\n" +
		"page=%2Fdocente%2Fturma%2Fturma.html%3Faction%3Dlist\n" +
		"scriptSessionId=" + scriptSessionId + "/MysIgVp-EK2*u7G*8\n" // Where did this come from ??????

	postBody := strings.NewReader(bodyText)

	req, err := http.NewRequest("POST", "https://portal.ufsm.br/docente/dwr/call/plaincall/gradeHorariosAjaxService.horarios.dwr", postBody)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error creating Discipline request: "+err.Error())
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Content-Type", "text/plain")

	resp, err := utils.Client.Do(req)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making Discipline request: "+err.Error())
		return
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading Discipline request "+err.Error())
	}

	sessionId := ParseSessionId(responseBody)

	if sessionId == "" {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError,
			"Error parsing SessionId: SessionId not found in "+string(responseBody))
		return
	}

}

func HandleTeacherDisplicines(w http.ResponseWriter, r *http.Request) {
	response := TeacherDisciplinesResponse{
		Disciplines: []TeacherDiscipline{
			{Id: "1", Code: "ELC1001", Name: "Cálculo I", Time: "8:00am"},
			{Id: "2", Code: "ELC1002", Name: "Álgebra Linear", Time: "10:30am"},
			{Id: "3", Code: "ELC1003", Name: "Física I", Time: "2:00pm"},
			{Id: "4", Code: "ELC1004", Name: "Programação I", Time: "4:30pm"},
		},
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
