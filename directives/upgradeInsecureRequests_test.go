// Directives for Content-Security-Policy headers.
package directives

import (
	"testing"
)

func TestUpgradeInsecureRequests(t *testing.T) {
	gotVal := UpgradeInsecureRequests()
	wantVal := "upgrade-insecure-requests;"

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
