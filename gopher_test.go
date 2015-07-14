package gopher

import "testing"

func TestNewApp(t *testing.T) {
	Config()
	if Render == nil {
		t.Error("Expected to get a pointer to Render, but got nil instead")
	}
}
