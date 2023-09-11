package src

import "testing"

func TestEmptyHost(t *testing.T) {
	host := Host("   ")
	gotVal := host.String()

	if gotVal != "" {
		t.Errorf("got: %s, want: empty string", gotVal)
	}
}

func TestInvalidHost(t *testing.T) {
	host := Host("^example.com  / api)/,3")
	gotVal := host.String()

	if gotVal != "" {
		t.Errorf("got: %s, want: empty string", gotVal)
	}
}

func TestWildcardHostWithScheme(t *testing.T) {
	wantVal := "http://*.example.com"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestHost(t *testing.T) {
	wantVal := "mail.example.com:443"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestHostWithScheme(t *testing.T) {
	wantVal := "https://store.example.com"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestWildcardHost(t *testing.T) {
	wantVal := "*.example.com"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestWildcardHostWithSchemeAndPath(t *testing.T) {
	wantVal := "https://*.example.com:12/path/to/file.js"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestWebsocketHost(t *testing.T) {
	wantVal := "ws://example.com"

	host := Host(wantVal)
	gotVal := host.String()

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
