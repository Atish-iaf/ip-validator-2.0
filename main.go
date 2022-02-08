package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// validateIPv4 function to check a given string is valid ipv4 or not.
func validateIPv4(line string) (bool, error) {
	ipv4, err := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)
	if err != nil {
		return false, err
	}
	if ipv4.MatchString(line) {
		return true, err
	} else {
		return false, err
	}
}

// validateIPv6 function to check a given string is valid ipv6 or not.
func validateIPv6(line string) (bool, error) {
	ipv6, err := regexp.Compile(`^((([0-9a-fA-F]){1,4})\:){7}([0-9a-fA-F]){1,4}$`)
	if err != nil {
		return false, err
	}
	if ipv6.MatchString(line) {
		return true, err
	} else {
		return false, err
	}
}

func ip(w http.ResponseWriter, r *http.Request) {

	inputIP := r.FormValue("ip")

	if inputIP == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
	}

	ipv4, ipv4err := validateIPv4(inputIP)
	ipv6, ipv6err := validateIPv6(inputIP)

	if ipv4err != nil {
		http.Error(w, ipv4err.Error(), http.StatusBadRequest)
	} else if ipv4 {
		fmt.Fprintln(w, "validIPv4")
	} else if ipv6err != nil {
		http.Error(w, ipv6err.Error(), http.StatusBadRequest)
	} else if ipv6 {
		fmt.Fprintln(w, "validIPv6")
	} else {
		fmt.Fprintln(w, "invalidIP")
	}

}

func main() {

	http.HandleFunc("/", ip)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
