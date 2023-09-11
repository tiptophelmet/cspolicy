package cspolicy

import (
	"testing"

	"github.com/tiptophelmet/cspolicy/directives"
	"github.com/tiptophelmet/cspolicy/src"
)

func TestEmptyBuild(t *testing.T) {
	gotVal := Build()
	wantVal := ""

	if gotVal != wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}

func TestBuild(t *testing.T) {
	gotVal := Build(
		directives.DefaultSrc(src.None()),
		directives.BaseURI(src.Self(), src.Host("*.example.com")),
		directives.ChildSrc(
			src.Host("cdn.example.com/assets"),
			src.Host("resources.example.com/artifacts"),
		),
		directives.ConnectSrc(
			src.Host("uploads.example.com"),
			src.Host("status.example.com"),
			src.Host("api.example.com"),
		),
		directives.FrameSrc(
			src.Host("notes.example.com"),
			src.Host("viewbox.example.com"),
		),
		directives.ImgSrc(
			src.Self(),
			src.Scheme("data:"),
			src.Host("media.example.com"),
			src.Host("avatars.example.com"),
		),
		directives.UpgradeInsecureRequests(),
	)

	wantVal := "default-src 'none'; base-uri 'self' *.example.com; " +
		"child-src cdn.example.com/assets resources.example.com/artifacts; " +
		"connect-src uploads.example.com status.example.com api.example.com; " +
		"frame-src notes.example.com viewbox.example.com; " +
		"img-src 'self' data: media.example.com avatars.example.com; " +
		"upgrade-insecure-requests"

	if gotVal == wantVal {
		t.Errorf("got: %s, want: %s", gotVal, wantVal)
	}
}
