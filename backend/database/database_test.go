package database

import (
	"log"
	"testing"
)

func TestQuery(t *testing.T) {

	err := OpenDB("../database.db")
	if err != nil {
		log.Fatal("Error Initializing Database")
	}
	err = RunMigrations()
	if err != nil {
		log.Fatal("Error Migrating Database")
	}

	s, err := FetchStudentinfo("980485")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v", s)

}
