package api

import "net/http"

func CheckMethodMW(method string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				t := getCType(r)
				headers := getHeaders(t)
				WriteError(w, headers, ErrMethodNotAllowed)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
