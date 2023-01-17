package webr

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
)

func viewError(err error) error {
	return errors.WithStack(err)
}

var RenderingError = renderingError{}

type renderingError struct {
	err error
}

func (e renderingError) Error() string {
	return "rendering error: " + e.err.Error()
}
func (e renderingError) Unwrap() error {
	return e.err
}
func (e renderingError) Is(target error) bool {
	if _, ok := target.(renderingError); ok {
		return true
	}
	return false
}

func render(w http.ResponseWriter, r *http.Request, tmpl *template.Template, data interface{}) error {
	if err := tmpl.Execute(w, data); err != nil {
		return renderingError{errors.WithStack(err)}
	}
	return nil
}

type ErrorModel struct {
	Error error
}

// View is a function that handles a request.
// It implements http.Handler, and can thus be used as a handler directly.
// All views should use this type.
//
// If the view returns an error, it will be handled by the ViewHandler.
type View func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP makes View implement http.Handler.
func (v View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v(w, r); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		panic(err)
	}
}
