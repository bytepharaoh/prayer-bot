package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDuration_MarshalUnmarshalJSON(t *testing.T) {
	d := Duration(90 * time.Minute)

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("marshal duration: %v", err)
	}

	if got := string(b); got != "\"1h30m0s\"" {
		t.Fatalf("unexpected marshal %s", got)
	}

	var decoded Duration
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal duration: %v", err)
	}

	if time.Duration(decoded) != 90*time.Minute {
		t.Fatalf("decoded duration mismatch: want %v got %v", 90*time.Minute, decoded.Duration())
	}
}

func TestDuration_UnmarshalJSON_Invalid(t *testing.T) {
	var d Duration
	err := json.Unmarshal([]byte(`"not valid"`), &d)
	if err == nil {
		t.Fatalf("expected error for invalid duration")
	}
}
