package main

import "net/http"

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Морковь"))
	})

	http.ListenAndServe(":5000", nil)
}
