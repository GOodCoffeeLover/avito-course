package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		host, err := os.Hostname()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Can't get who am i: %v", err)))
			log.Printf("Error: %v", err)
		}
		resp := fmt.Sprintf("I'm %v\n", host)
		log.Printf("Sending ok: %v", resp)
		w.Write([]byte(resp))
	})

	log.Println("Starting http server ...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Can't start: %v", err)
	}
}
