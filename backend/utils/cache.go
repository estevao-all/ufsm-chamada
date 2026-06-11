package utils

import (
	"net/http"
	"time"

	go_cache "github.com/patrickmn/go-cache"
)

type responseRecorder struct {
	http.ResponseWriter
	Body       []byte
	StatusCode int
}

func (r *responseRecorder) Write(bytes []byte) (int, error) {
	r.Body = append(r.Body, bytes...)
	return r.ResponseWriter.Write(bytes)
}

func (r *responseRecorder) WriteHeader(status_code int) {
	r.StatusCode = status_code
	r.ResponseWriter.WriteHeader(status_code)
}

type cacheEntry struct {
	Body       []byte
	Headers    http.Header
	StatusCode int
}

var api_cache = go_cache.New(1*time.Hour, 1*time.Hour)

func cacheKey(r *http.Request) string {
	session_cookie, err := r.Cookie("JSESSIONIDSSO")
	if err != nil {
		return r.URL.String()
	}
	return r.URL.String() + "|" + session_cookie.Value
}

func WithCache(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			handler(w, r)
			return
		}

		cache_key := cacheKey(r)

		if cached, found := api_cache.Get(cache_key); found {
			entry := cached.(cacheEntry)
			for key, values := range entry.Headers {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}

			w.WriteHeader(entry.StatusCode)
			w.Write(entry.Body)

			return
		}

		rec := &responseRecorder{ResponseWriter: w, StatusCode: http.StatusOK}
		handler(rec, r)

		api_cache.Set(cache_key, cacheEntry{
			Body:       rec.Body,
			Headers:    w.Header().Clone(),
			StatusCode: rec.StatusCode,
		}, go_cache.DefaultExpiration)
	}
}
