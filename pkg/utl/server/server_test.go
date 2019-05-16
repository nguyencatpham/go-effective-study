package server_test

import (
	"testing"

	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/server"
)

// Improve tests
func TestNew(t *testing.T) {
	e := server.New()
	if e == nil {
		t.Errorf("Server should not be nil")
	}
}
