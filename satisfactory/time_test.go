package satisfactory

import (
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	var ct CustomTime
	jsonStr := `"2024.10.04-23.24.38"`

	err := ct.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}

	expectedTime, _ := time.Parse("2006.01.02-15.04.05", "2024.10.04-23.24.38")
	if !ct.Time.Equal(expectedTime) {
		t.Errorf("Expected time %v, got %v", expectedTime, ct.Time)
	}
}
