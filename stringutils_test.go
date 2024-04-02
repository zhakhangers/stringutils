// stringutils_test.go
package stringutils

import (
	"fmt"
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

func TestRuneCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty string", "", 0},
		{"ASCII characters", "hello", 5},
		{"Unicode characters", "こんにちは", 5},
		{"Mixed characters", "hello こんにちは", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringutils.RuneCount(tt.input)
			if got != tt.want {
				t.Errorf("RuneCount(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func ExampleReverse() {
	fmt.Println(stringutils.Reverse("こんにちは"))
	// Output: はちにんこ
}

func ExampleRuneCount() {
	fmt.Println(stringutils.RuneCount("hello"))
	// Output: 5
}
