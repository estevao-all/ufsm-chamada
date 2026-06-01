package routes

import (
	"strings"
	"testing"
)

func TestParseSessionId(t *testing.T) {

	testData := []byte(`throw 'allowScriptTagRemoting is false.';\n//#DWR-REPLY\n//#DWR-START#\n(function(){\r\nif(!window.dwr)return;\r\nvar dwr=window.dwr._[1];\ndwr.engine.remote.handleCallback(\"0\",\"0\",\"YTYvGc~Ngqt81Iu0C53vM!YU!lHzrdsIgVp\");\n})();\n//#DWR-END#\n`)
	r := ParseSessionId(testData)
	t.Logf("%v", r)
}

func TestUrlEncoding(t *testing.T) {
	scriptSessionId := "YTYvGc~Ngqt81Iu0C53vM!YU!lHzrdsIgVp"

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
		"scriptSessionId=" + scriptSessionId + "/MysIgVp-EK2*u7G*8\n"

	postBody := strings.NewReader(bodyText)
	t.Log(postBody)

}
