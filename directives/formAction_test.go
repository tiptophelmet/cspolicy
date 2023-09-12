// Directives for Content-Security-Policy headers.
package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyFormAction(t *testing.T) {
	gotVal := FormAction()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFormAction(t *testing.T) {
	gotVal := FormAction(
		src.Host("api.example.com"),
		src.Host("collect.example.com"),
	)

	wantVal := "form-action api.example.com collect.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
