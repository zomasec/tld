# tld v1.0
The `tld` package provides functionality to parse URLs and extract various components such as subdomain, domain, top-level domain (TLD), and port. It is particularly useful for scenarios where you need to analyze or manipulate URLs in your Go applications.

# tld

[![Go Reference](https://pkg.go.dev/badge/github.com/zomasec/tld.svg)](https://pkg.go.dev/github.com/zomasec/tld)

A Go package for parsing URLs and extracting subdomain, domain, top-level domain (TLD), and port information.

## Description

The `tld` package provides functionality to parse URLs and extract various components such as subdomain, domain, top-level domain (TLD), and port. It is particularly useful for scenarios where you need to analyze or manipulate URLs in your Go applications.

### Features

- Parse URLs and extract subdomain, domain, TLD, and port information.
- Handles multiple subdomains extraction.
- Handle edge cases gracefully.
- Error handling for robustness.
- Compatible with Go modules.

## Installation

To use the `tld` package in your Go project, you can simply import it via Go modules:

```bash
go get github.com/zomasec/tld
```

## Usage
Here's a basic example demonstrating how to use the tld package:

```go
package main

import (
	"fmt"
	"github.com/zomasec/tld"
)

func main() {
	urlString := "https://subdomain1.subdomain2.example.com:8080/path/to/resource"
	parsedURL, err := tld.Parse(urlString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Subdomains:", parsedURL.Subdomains)
	fmt.Println("Domain:", parsedURL.Domain)
	fmt.Println("TLD:", parsedURL.TLD)
	fmt.Println("Port:", parsedURL.Port)
}

```

### Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or create a pull request on GitHub.
