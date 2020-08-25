package main

import (
	"testing"
	"github.com/kiddy_line_processor_test.git/database"
)

func Testdatabase(t *testing.T) {
	var want float32
	database.Insert_football(want)
	m := database.Select_sport("football")

	if m != want {
		t.Fatalf("Want %v, but got %v", want, m)
	}
}