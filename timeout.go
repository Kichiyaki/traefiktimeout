package traefiktimeout

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrTimeoutMustBePositive = errors.New("timeout must be a positive number")

// Config the plugin configuration.
type Config struct {
	Timeout string `json:"timeout"`
	Message string `json:"message,omitempty"`
}

type TraefikTimeout struct {
	name string
	next http.Handler
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Message: "Timeout",
	}
}

func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	timeout, err := time.ParseDuration(config.Timeout)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse given timeout: %w", err)
	}

	if timeout <= 0 {
		return nil, ErrTimeoutMustBePositive
	}

	return &TraefikTimeout{
		name: name,
		next: http.TimeoutHandler(next, timeout, config.Message),
	}, nil
}

func (mw *TraefikTimeout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw.next.ServeHTTP(w, r)
}
