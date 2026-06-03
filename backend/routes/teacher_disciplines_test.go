package routes

import (
	"strings"
	"testing"
)

func TestParseSessionId(t *testing.T) {

	testData := []byte(`throw 'allowScriptTagRemoting is false.';\n//#DWR-REPLY\n//#DWR-START#\n(function(){\r\nif(!window.dwr)return;\r\nvar dwr=window.dwr._[1];\ndwr.engine.remote.handleCallback("0","0","YTYvGc~Ngqt81Iu0C53vM!YU!lHzrdsIgVp");\n})();\n//#DWR-END#\n`)
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
		"c0-param0=\"string\":2026\n" +
		"c0-param1=\"string\":101\n" +
		"batchId=0\n" +
		"instanceId=0\n" +
		"page=%2Fdocente%2Fturma%2Fturma.html%3Faction%3Dlist\n" +
		"scriptSessionId=" + scriptSessionId + "/MysIgVp-EK2*u7G*8\n"

	postBody := strings.NewReader(bodyText)
	t.Log(postBody)

}

func TestParseDisciplines(t *testing.T) {

	testData := []byte(`"text": "throw 'allowScriptTagRemoting is false.';\n//#DWR-INSERT\n//#DWR-REPLY\n//#DWR-START#\n(function(){\r\nif(!window.dwr)return;\r\nvar dwr=window.dwr._[0];\ndwr.engine.remote.handleCallback(\"0\",\"0\",{ano:2026,horarios:{quarta:[{diaSemana:4,disciplinas:[\"INSTALA\\u00C7\\u00D5ES EL\\u00C9TRICAS RESIDENCIAIS E COMERCIAIS\"],fim:new Date(48600000),id:1105419,inicio:new Date(41400000),tipo:\"Te\\u00F3rica\",turmas:[{id:980485,nome:\"11_302 - Engenharia Civil\"}]}],quinta:[{diaSemana:5,disciplinas:[\"PROJETO INTEGRADOR EM ENGENHARIA EL\\u00C9TRICA II\"],fim:new Date(63000000),id:1103519,inicio:new Date(59400000),tipo:\"Te\\u00F3rica\",turmas:[{id:979000,nome:\"11_303 - Engenharia El\\u00E9trica\"}]},{diaSemana:5,disciplinas:[\"PROJETO INTEGRADOR EM ENGENHARIA EL\\u00C9TRICA II\"],fim:new Date(73800000),id:1103520,inicio:new Date(63000000),tipo:\"Pr\\u00E1tica\",turmas:[{id:979000,nome:\"11_303 - Engenharia El\\u00E9trica\"}]}],sabado:[],segunda:[{diaSemana:2,disciplinas:[\"INSTALA\\u00C7\\u00D5ES EL\\u00C9TRICAS RESIDENCIAIS E COMERCIAIS\"],fim:new Date(45000000),id:1105417,inicio:new Date(41400000),tipo:\"Te\\u00F3rica\",turmas:[{id:980485,nome:\"11_302 - Engenharia Civil\"}]},{diaSemana:2,disciplinas:[\"INSTALA\\u00C7\\u00D5ES EL\\u00C9TRICAS RESIDENCIAIS E COMERCIAIS\"],fim:new Date(48600000),id:1105418,inicio:new Date(45000000),tipo:\"Pr\\u00E1tica\",turmas:[{id:980485,nome:\"11_302 - Engenharia Civil\"}]}],sexta:[{diaSemana:6,disciplinas:[\"PROJETO INTEGRADOR EM ENGENHARIA DE COMPUTA\\u00C7\\u00C3O II\"],fim:new Date(63000000),id:1109060,inicio:new Date(59400000),tipo:\"Te\\u00F3rica\",turmas:[{id:983071,nome:\"10_312 - Curso de Engenharia de Computa\\u00E7\\u00E3o\"}]},{diaSemana:6,disciplinas:[\"PROJETO INTEGRADOR EM ENGENHARIA DE COMPUTA\\u00C7\\u00C3O II\"],fim:new Date(66600000),id:1109061,inicio:new Date(63000000),tipo:\"Pr\\u00E1tica extensionista\",turmas:[{id:983071,nome:\"10_312 - Curso de Engenharia de Computa\\u00E7\\u00E3o\"}]}],terca:[]},periodoItem:101});\n})();\n//#DWR-END#\n"`)

	s := ParseTeacherDisciplines(testData)

	r, err := parseSchedule(s)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%v", r.Horarios["quarta"])
}
