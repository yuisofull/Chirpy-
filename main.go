package main

import "net/http"

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsMux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
