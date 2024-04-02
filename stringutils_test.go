// stringutils_test.go
package stringutils_test

import (
	"testing"

	"github.com/zhakhangers/stringutils"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name, input, want string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"multiple characters", "abc", "cba"},
		{"unicode characters", "こんにちは", "はちにんこ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringutils.Reverse(tt.input)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}

}
