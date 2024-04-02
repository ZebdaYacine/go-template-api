package utils

import "testing"

func TestGenerateNewCode(t *testing.T) {
	got := GenerateNewCode(853)
	t.Errorf("GenerateNewCode() = %v, want %v", got, 1)
}
