package godomains

import (
	"log"
	"net"
)

// IsAvailable checks if a domain is not taken by performing dns lookups.
func IsAvailable(domain string) bool {
	_, err := net.LookupHost(domain)
	return err != nil
}

func checkAvailability(domain string) {
	if IsAvailable(domain) {
		log.Println(domain + " is available")
	} else {
		log.Println(domain + " is taken")
	}
}
