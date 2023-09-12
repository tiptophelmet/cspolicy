package directives

// Constructs the CSP 'upgrade-insecure-requests' directive.
// The 'upgrade-insecure-requests' directive instructs user agents to treat all of a site's insecure URLs as though they have been replaced with secure URLs.
// This directive is intended for web sites with large numbers of insecure legacy URLs that need to be rewritten.
// See: [CSP.Directives.upgrade-insecure-requests].
//
// [CSP.Directives.upgrade-insecure-requests]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/upgrade-insecure-requests
func UpgradeInsecureRequests() string {
	return "upgrade-insecure-requests;"
}
