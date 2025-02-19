package web

// MidHandler is a handler function designed to run code before and/or after
// another Handler. It is designed to remove boilerplate or other concerns not
// direct to any given app Handler.
type MidHandler func(Handler) Handler

// wrapMiddleware creates a new handler by wrapping middleware around a final handler.
// The middleware's Handlers will be executed by requests in the order they are provided.
func wrapMiddleware(mv []MidHandler, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mv) - 1; i >= 0; i-- {
		mwFunc := mv[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}

	return handler
}
