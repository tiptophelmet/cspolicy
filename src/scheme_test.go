// Sources for Content-Security-Policy directives.
package src

import (
	"testing"
)

func TestDataScheme(t *testing.T) {
	wantVal := "data:"
	gotVal := Scheme(wantVal).String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestMediastreamScheme(t *testing.T) {
	wantVal := "mediastream:"
	gotVal := Scheme(wantVal).String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestBlobScheme(t *testing.T) {
	wantVal := "blob:"
	gotVal := Scheme(wantVal).String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFileSystemScheme(t *testing.T) {
	wantVal := "filesystem:"
	gotVal := Scheme(wantVal).String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
