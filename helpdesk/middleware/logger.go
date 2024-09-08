package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrapWriter struct {
	http.ResponseWriter
	statusCode int
}

// modify the header for print status code in logger
func (w *wrapWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reponseWtriter := &wrapWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(reponseWtriter, r)
		log.Printf("%v %v %v %v\n", reponseWtriter.statusCode, r.URL.Path, r.Method, time.Since(start))
	})
}
