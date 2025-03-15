package web

import "errors"

// shutdownError is a type used to help with the graceful termination of the service.
type shutdownError struct {
	Message string
}

// Error is the implementation of the error interface.
func (se *shutdownError) Error() string {
	return se.Message
}

// NewShutdownError returns an error that causes the framework to signal a graceful shutdown.
func NewShutdownError(message string) error {
	return &shutdownError{message}
}

func IsShutdown(err error) bool {
	var se *shutdownError
	return errors.As(err, &se)
}
