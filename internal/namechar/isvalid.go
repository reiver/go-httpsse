package namechar

// isValid return whether the rune is a valid "name-char" character as defined by the HTTP-SSE specification:
// https://html.spec.whatwg.org/multipage/server-sent-events.html
//
//	name-char     = %x0000-0009 / %x000B-000C / %x000E-0039 / %x003B-10FFFF
//	                ; a scalar value other than U+000A LINE FEED (LF), U+000D CARRIAGE RETURN (CR), or U+003A
func isValid(r rune) bool {
	switch {
	case '\u0000' <= r && r <= '\u0009':
		return true
	case '\u000B' <= r && r <= '\u000C':
		return true
	case '\u000E' <= r && r <= '\u0039':
		return true
	case '\u003B' <= r && r <= '\U0010FFFF':
		return true
	default:
		return false
	}
}
