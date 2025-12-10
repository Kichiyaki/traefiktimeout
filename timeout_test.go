package traefiktimeout_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Kichiyaki/traefiktimeout"
)

func TestTraefikTimeout(t *testing.T) {
	t.Parallel()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := time.NewTimer(time.Second)
		defer timer.Stop()

		select {
		case <-r.Context().Done():
			return
		case <-timer.C:
			w.WriteHeader(http.StatusGatewayTimeout)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		t.Parallel()

		cfg := traefiktimeout.CreateConfig()

		cfg.Timeout = time.Nanosecond.String()

		h, err := traefiktimeout.New(context.Background(), next, cfg, "traefiktimeout")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		h.ServeHTTP(rr, req.WithContext(context.Background()))

		if rr.Code != http.StatusServiceUnavailable {
			t.Errorf("got status code %d, want %d", rr.Code, http.StatusServiceUnavailable)
		}
	})

	t.Run("parent context canceled", func(t *testing.T) {
		t.Parallel()

		cfg := traefiktimeout.CreateConfig()

		cfg.Timeout = time.Second.String()

		h, err := traefiktimeout.New(context.Background(), next, cfg, "traefiktimeout")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		h.ServeHTTP(rr, req.WithContext(ctx))

		if rr.Code != http.StatusServiceUnavailable {
			t.Errorf("got status code %d, want %d", rr.Code, http.StatusServiceUnavailable)
		}
	})
}
