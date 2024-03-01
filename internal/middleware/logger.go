package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

type UserLog struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (r customResponseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r customResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
}

func Logger(log *zerolog.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := &customResponseWriter{
				ResponseWriter: w,
				body:           &bytes.Buffer{},
				statusCode:     http.StatusOK,
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error reading body: %v", err)
			}

			var payload map[string]interface{}
			if len(body) > 0 {
				json.Unmarshal(body, &payload)
			}

			defer func() {
				status := ww.statusCode
				log.Info().Timestamp().Fields(map[string]interface{}{
					"@path":   r.URL.Path,
					"@time":   time.Now().Format(time.RFC3339),
					"@method": r.Method,
					"@status": status,
					"@body":   payload,
				}).Msg("Handler Request")

			}()

			newBody := io.NopCloser(bytes.NewBuffer(body))
			r.Body = newBody
			h.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
