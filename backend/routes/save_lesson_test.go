package routes

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestUrlEncoded(t *testing.T) {
	req_body := url.Values{}

	t.Log(time.Now().Format("02/01/2006"))

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

	req_body.Set("save", "")

	t.Log(strings.NewReader(req_body.Encode()))

}
