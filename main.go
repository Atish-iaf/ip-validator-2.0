package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// validateIPv4 function to check a given string is valid ipv4 or not.
func validateIPv4(line string) bool {
	ipv4, err := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)
	if err != nil {
		log.Fatal(err)
	}
	if ipv4.MatchString(line) {
		return true
	} else {
		return false
	}
}

// validateIPv6 function to check a given string is valid ipv6 or not.
func validateIPv6(line string) bool {
	ipv6, err := regexp.Compile(`^((([0-9a-fA-F]){1,4})\:){7}([0-9a-fA-F]){1,4}$`)
	if err != nil {
		log.Fatal(err)
	}
	if ipv6.MatchString(line) {
		return true
	} else {
		return false
	}
}

func ip(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		inputIP := r.FormValue("IP")
		if validateIPv4(inputIP) {
			w.Write([]byte(`<div style = "display:flex; justify-content:center; align-items:center; height:100vh;"><h2>valid IPv4 </h2></div>`))
		} else if validateIPv6(inputIP) {
			w.Write([]byte(`<div style = "display:flex; justify-content:center; align-items:center; height:100vh;"><h2>valid IPv6 </h2></div>`))
		} else {
			w.Write([]byte(`<div style = "display:flex; justify-content:center; align-items:center; height:100vh;"><h2>invalid IP </h2></div>`))
		}
	default:
		w.Write([]byte(`<div style = "display:flex; justify-content:center; align-items:center; height:100vh;"><h2>Only GET and POST methods are supported. </h2></div>`))
	}
}

func main() {

	http.HandleFunc("/", ip)

	fmt.Println("Starting server for HTTP POST...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
