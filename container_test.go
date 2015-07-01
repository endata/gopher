package gopher

import (
	"testing"

	"github.com/gopherlabs/gopher"
)

func TestNewApp(t *testing.T) {
	app := gopher.App()
	if app == nil {
		t.Error("Expected to get a pointer to the container, but got nil instead")
	}
}
