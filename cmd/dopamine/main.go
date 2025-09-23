package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Морковь"))
	})

	fmt.Print("Server is started")

	http.ListenAndServe(":8000", nil)
}
