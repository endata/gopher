package gopher

import "testing"

func TestNewApp(t *testing.T) {
	app := Config()
	if app == nil {
		t.Error("Expected to get a pointer to the container, but got nil instead")
	}
}
