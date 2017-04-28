package fake

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test #1", "Hello World from server"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); strings.Index(got, tt.want) != 0 {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
