package routes

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Schedule struct {
	Ano         int                     `json:"ano"`
	Horarios    map[string][]DayClasses `json:"horarios"`
	PeriodoItem int                     `json:"periodoItem"`
}

type Class struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type DayClasses struct {
	DiaSemana   int      `json:"diaSemana"`
	Disciplinas []string `json:"disciplinas"`
	Fim         string   `json:"fim"`
	ID          int      `json:"id"`
	Inicio      string   `json:"inicio"`
	Tipo        string   `json:"tipo"`
	Turmas      []Class  `json:"turmas"`
}

func formatScheduleTime(ms int64) string {
	totalSeconds := ms / 1000
	hours := totalSeconds/3600 - 3 // Timezone offset
	minutes := (totalSeconds % 3600) / 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Custom unmarshaler for millisecond timestamps
func (c *DayClasses) UnmarshalJSON(data []byte) error {
	type Alias DayClasses
	aux := &struct {
		Fim    int64 `json:"fim"`
		Inicio int64 `json:"inicio"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Fim = formatScheduleTime(aux.Fim)
	c.Inicio = formatScheduleTime(aux.Inicio)
	return nil
}

func parseSchedule(data string) (*Schedule, error) {

	slashRe := regexp.MustCompile(`\\\\`)
	cleaned := slashRe.ReplaceAllString(data, `\`)

	slashQuoteRe := regexp.MustCompile(`\\"`)
	cleaned = slashQuoteRe.ReplaceAllString(cleaned, `"`)

	keyRe := regexp.MustCompile(`(\w+):`)
	cleaned = keyRe.ReplaceAllString(cleaned, `"$1":`)

	// Strip new Date(...) wrappers, leaving just the millisecond values
	re := regexp.MustCompile(`new Date\((\d+)\)`)
	cleaned = re.ReplaceAllString(cleaned, `$1`)

	var schedule Schedule
	err := json.Unmarshal([]byte(cleaned), &schedule)
	return &schedule, err
}
