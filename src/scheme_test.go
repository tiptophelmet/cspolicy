package src

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src/scheme"
)

func TestDataScheme(t *testing.T) {
	gotVal := Scheme(scheme.Data()).String()
	wantVal := "data:"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestMediastreamScheme(t *testing.T) {
	gotVal := Scheme(scheme.Mediastream()).String()
	wantVal := "mediastream:"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestBlobScheme(t *testing.T) {
	gotVal := Scheme(scheme.Blob()).String()
	wantVal := "blob:"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestFileSystemScheme(t *testing.T) {
	gotVal := Scheme(scheme.FileSystem()).String()
	wantVal := "filesystem:"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
