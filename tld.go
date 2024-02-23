package tld

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// URL embeds net/url.URL and adds additional fields for subdomains, domain, TLD, and port.
type URL struct {
	*url.URL
	Subdomains []string // Subdomains of the URL.
	Domain     string   // Domain of the URL.
	TLD        string   // Top-level domain of the URL.
	Port       string   // Port of the URL.
}

// Parse parses the input URL string and returns a tld.URL, which contains extra fields for subdomains, domain, TLD, and port.
func Parse(s string) (*URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if u.Host == "" {
		return &URL{URL: u}, nil
	}

	subdomains, domain, port := parseHost(u.Host)
	etldPlusOne, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return nil, fmt.Errorf("failed to determine effective TLD+1: %w", err)
	}

	return &URL{
		URL:        u,
		Subdomains: subdomains,
		Domain:     domain,
		TLD:        strings.TrimPrefix(etldPlusOne, domain+"."),
		Port:       port,
	}, nil
}

// parseHost parses the host part of a URL and returns the subdomains, domain, and port.
func parseHost(host string) ([]string, string, string) {
	var subdomains []string
	var domain, port string

	parts := strings.Split(host, ":")
	if len(parts) > 1 {
		host = parts[0] // Extract host without port
		port = parts[1]
	}

	subdomainParts := strings.Split(host, ".")
	for i := len(subdomainParts) - 1; i >= 0; i-- {
		if subdomainParts[i] != "" {
			subdomains = append(subdomains, subdomainParts[i])
		}
	}

	if len(subdomains) > 0 {
		domain = subdomains[0]
		if len(subdomains) > 1 {
			domain = subdomains[1] + "." + domain
			subdomains = subdomains[2:]
		} else {
			subdomains = nil
		}
	}
	return subdomains, domain, port
}
