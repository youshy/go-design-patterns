package main

import (
	"fmt"
	"log"
	"net/http"
)

// not using OAuth2 or JWT for readability
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Checking if Authorized...\n")

		if val, ok := r.Header["Authorized"]; ok {
			log.Println(val)
			if val[0] == "true" {
				log.Printf("Correct header!")
				endpoint(w, r)
			}
		} else {
			log.Printf("Unauthorized access")
			fmt.Fprintf(w, "Unauthorized access")
		}
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint: homePage\n")
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8081", nil))
	log.Printf("Serving on :8081\n")
}

func main() {
	handleRequests()
}
