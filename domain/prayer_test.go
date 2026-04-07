package domain

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"minutes only", 30 * time.Minute, "30m"},
		{"hours and minutes", 90 * time.Minute, "1h30m"},
		{"exact hours", 2 * time.Hour, "2h0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.d); got != tt.want {
				t.Fatalf("FormatDuration(%v) = %s, want %s", tt.d, got, tt.want)
			}
		})
	}
}

func TestDateUTC(t *testing.T) {
	got := DateUTC(31, time.December, 2024)
	if got.Year() != 2024 || got.Month() != time.December || got.Day() != 31 {
		t.Fatalf("unexpected date: %v", got)
	}
	if got.Location() != time.UTC {
		t.Fatalf("expected UTC location")
	}
	if got.Hour() != 0 || got.Minute() != 0 || got.Second() != 0 {
		t.Fatalf("expected midnight UTC, got %v", got)
	}
}
