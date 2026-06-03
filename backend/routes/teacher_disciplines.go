package routes

import (
	"backend/utils"
	"errors"
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

var sessionIdRegex = regexp.MustCompile(`(?s)handleCallback\(\\"0\\",\\"0\\",\\"(.+?)\\"\)`)

func ParseSessionId(responseBody []byte) string {
	sessionId := sessionIdRegex.FindStringSubmatch(string(responseBody))
	if len(sessionId) < 2 {
		return ""
	}
	return sessionId[1]
}

func GetScriptSessionId(r *http.Request) (string, error) {

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
		return "", err
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Content-Type", "text/plain")

	resp, err := utils.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sessionId := ParseSessionId(responseBody)

	if sessionId == "" {
		return "", errors.New("Error parsing SessionId: SessionId not found in " + string(responseBody))
	}

	return sessionId, nil
}

var DisciplineRegex = regexp.MustCompile(`{.+?({ano:\d{4},horarios:{.+},periodoItem:\d{3}})`)

func ParseTeacherDisciplines(responseBody []byte) string {
	match := DisciplineRegex.FindStringSubmatch(string(responseBody))
	if len(match) < 2 {
		return ""
	}

	return match[1]
}

func GetDisciplines(r *http.Request, scriptSessionId string) (*Schedule, error) {

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
		return nil, err
	}

	req.Header.Set("Host", UFSM_PORTAL_HOST)
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Content-Type", "text/plain")

	resp, err := utils.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	schedule, err := parseSchedule(string(responseBody))
	if err != nil {
		return nil, err
	}

	return schedule, nil

}

func HandleTeacherDisplicines(w http.ResponseWriter, r *http.Request) {
	scriptSessionId, err := GetScriptSessionId(r)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making ScriptSessionId request "+err.Error())
	}

	schedule, err := GetDisciplines(r, scriptSessionId)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error making Schedule request "+err.Error())
	}

	utils.WriteJSON(w, http.StatusOK, *schedule)
}
