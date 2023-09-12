// Sources for Content-Security-Policy directives.
package src

// Constraint to prevent random [sources] for [CSP directives].
//
// [sources]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/Sources
// [CSP directives]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy#directives
type SourceVal interface {
	String() string
}
