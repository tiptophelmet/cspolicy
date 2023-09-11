package directives

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyConnectSrc(t *testing.T) {
	gotVal := ConnectSrc()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestConnectSrc(t *testing.T) {
	gotVal := ConnectSrc(
		src.Host("uploads.example.com"),
		src.Host("status.example.com"),
		src.Host("api.example.com"),
	)

	wantVal := "connect-src uploads.example.com status.example.com api.example.com;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
