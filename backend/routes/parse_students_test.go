package routes

import (
	"os"
	"testing"
)

func TestStudentParse(t *testing.T) {
	data, err := os.ReadFile(`..\..\..\mock_student_data.txt`)
	if err != nil {
		t.Logf("read error: %v", err)
		return
	}

	s := string(data)
	info, err := ParseStudents(s)
	if err != nil {
		t.Logf("read error: %v", err)
		return
	}
	t.Logf("%v", info)
}
