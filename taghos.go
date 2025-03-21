package fasthttp

// Query returns the first value for the named component of the URL query parameters.
// If key is not present, it returns the specified default value or an empty string.
func (ctx *RequestCtx) Query(name string, defaultValue ...string) string {
	if v := ctx.QueryArgs().Peek(name); v != nil {
		return string(v)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}
