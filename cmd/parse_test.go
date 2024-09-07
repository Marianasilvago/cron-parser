package cmd

import (
	"testing"
)

// TestExpandField tests various field expansion cases.
func TestExpandField(t *testing.T) {
	tests := []struct {
		field string
		min   int
		max   int
		want  string
	}{
		{"*", 0, 59, "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59"}, // Full range
		{"*/15", 0, 59, "0 15 30 45"},        // Step of 15
		{"1,15", 0, 59, "1 15"},              // Specific minutes
		{"1-5", 1, 5, "1 2 3 4 5"},           // Range expansion
		{"5", 0, 59, "5"},                    // Single value
		{"*/5", 1, 31, "1 6 11 16 21 26 31"}, // Step of 5 for day of month
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := expandField(tt.field, tt.min, tt.max)
			if got != tt.want {
				t.Errorf("expandField(%q, %d, %d) = %q; want %q", tt.field, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
