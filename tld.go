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
	Subdomains []string
	Domain     string
	TLD        string
	Port       string
}

// Parse parses the input URL string or domain and returns a tld.URL, which contains extra fields for subdomains, domain, TLD, and port.
func Parse(s string) (*URL, error) {
	var u *url.URL
	var err error

	// Check if the input contains scheme (http:// or https://)
	if strings.Contains(s, "://") {
		u, err = url.Parse(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}
	} else {
		u = &url.URL{Host: s}
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

func extractSubdomains(fullDomain, etldPlusOne string) []string {
	var subdomains []string
	if fullDomain != etldPlusOne {
		subdomainParts := strings.Split(fullDomain, ".")
		for i := 0; i < len(subdomainParts)-2; i++ {
			subdomains = append(subdomains, subdomainParts[i])
		}
	}
	return subdomains
}

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
