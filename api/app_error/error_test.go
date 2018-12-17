package app_error

import (
	"fmt"
	"testing"
)

func TestAppErrorIsError(t *testing.T) {
	var err error
	err = New(fmt.Errorf("I'M A TEAPOT"), 418)

	if err.Error() != "I'M A TEAPOT" {
		t.Fatalf("err.Error() = %s, want %s", err.Error(), "I'M A TEAPOT")
	}
}
