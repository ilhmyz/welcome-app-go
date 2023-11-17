package main

import (
	"fmt"
	"log"
	"net/http"
)

func namaHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Path[len("/welcome/"):]
	if len(nama) > 0 {
		fmt.Fprintf(w, "Selamat datang %s", nama)
		log.Printf(nama)
	} else {
		fmt.Fprint(w, "Anonymous")
		log.Printf("Anonymous")
	}
}

func main() {
	http.HandleFunc("/welcome/", namaHandler)

	fmt.Println("App berjalan pada port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
