// Package main provides functionality for parsing URLs and extracting subdomains, domain, TLD, and port.
package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// URL embeds net/url.URL and adds additional fields for subdomains, domain, TLD, and port.
type URL struct {
	*url.URL
	// Subdomains holds the subdomains extracted from the URL.
	Subdomains []string
	// Domain holds the main domain extracted from the URL.
	Domain string
	// TLD holds the top-level domain extracted from the URL.
	TLD string
	// Port holds the port number extracted from the URL, if present.
	Port string
}

// Parse parses the input URL string and returns a tld.URL, which contains extra fields for subdomains, domain, TLD, and port.
func Parse(s string) (*URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	if u.Host == "" {
		return &URL{URL: u}, nil
	}

	domain, port := parseHost(u.Host)
	etldPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return nil, fmt.Errorf("failed to determine effective TLD+1: %v", err)
	}

	subdomains := extractSubdomains(domain, etldPlusOne)

	i := strings.Index(etldPlusOne, ".")
	if i < 0 {
		return nil, fmt.Errorf("failed to extract domain and TLD from URL: %s", s)
	}
	domainName := etldPlusOne[:i]
	tld := etldPlusOne[i+1:]

	return &URL{
		URL:        u,
		Subdomains: subdomains,
		Domain:     domainName,
		TLD:        tld,
		Port:       port,
	}, nil
}

// extractSubdomains extracts subdomains from the full domain name.
func extractSubdomains(fullDomain, etldPlusOne string) []string {
	var subdomains []string
	if fullDomain != etldPlusOne {
		subdomainParts := strings.Split(fullDomain, ".")
		for i := 0; i < len(subdomainParts)-3; i++ {
			subdomains = append(subdomains, subdomainParts[i])
		}
	}
	return subdomains
}

// parseHost extracts the domain and port from the host part of a URL.
func parseHost(host string) (string, string) {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return host[:i], host[i+1:]
		} else if host[i] < '0' || host[i] > '9' {
			return host, ""
		}
	}
	return host, ""
}

func main() {
	urlString := "https://subdomain1.subdomain2.example.com.eg:8080/path/to/resource"
	parsedURL, err := Parse(urlString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Subdomains:", parsedURL.Subdomains)
	fmt.Println("Domain:", parsedURL.Domain)
	fmt.Println("TLD:", parsedURL.TLD)
	fmt.Println("Port:", parsedURL.Port)
}
