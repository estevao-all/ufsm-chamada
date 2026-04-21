package routes

import (
	"backend/utils"
	"net/http"
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
