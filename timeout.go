package traefiktimeout

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var ErrTimeoutMustBePositive = errors.New("timeout must be a positive number")

// Config the plugin configuration.
type Config struct {
	Timeout time.Duration `json:"timeout"`
	Message string        `json:"message,omitempty"`
}

type TraefikTimeout struct {
	name string
	cfg  *Config
	next http.Handler
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Message: "Timeout",
	}
}

func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Timeout <= 0 {
		return nil, ErrTimeoutMustBePositive
	}

	return &TraefikTimeout{
		name: name,
		cfg:  config,
		next: next,
	}, nil
}

func (mw *TraefikTimeout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.TimeoutHandler(mw.next, mw.cfg.Timeout, mw.cfg.Message).ServeHTTP(w, r)
}
