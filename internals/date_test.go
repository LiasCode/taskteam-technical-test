package internals

import (
	"testing"
	"time"
)

func TestDateCreation(t *testing.T) {
	if IsDateValue("2025-07-24", time.DateTime) {
		t.Errorf("Should throw")
	}

	newDate := CreateDateAsDateTime()

	if !IsDateValue(newDate, time.DateTime) {
		t.Errorf("Date should be in DateTime format")
	}
}
