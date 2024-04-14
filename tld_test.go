package tld

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("Case have one subdomain", func(t *testing.T) {
		want := "subdomain1"
		parsed, _ := Parse("https://subdomain1.example.com.eg")
		got := parsed.Subdomains[0]
		assert.Equalf(t, want, got, "want %s but got %s", want, got)
	})

	t.Run("Case have more than subdomain", func(t *testing.T) {
		want := []string{"subdomain1", "subdomain2", "subdomain3"}
		parsed, _ := Parse("subdomain1.subdomain2.subdomain3.example.com")
		got := parsed.Subdomains
		fmt.Println(got)
		assert.ElementsMatch(t, want, got, "want %s but got %s", want, got)

	})

	t.Run("Case i want tld", func(t *testing.T) {
		want := "com"
		parsed, _ := Parse("https://subdomain1.example.com")
		got := parsed.TLD
		assert.Equalf(t, want, got, "want %s but got %s", want, got)

	})

	t.Run("Case invalid url", func(t *testing.T) {
		_, err := Parse("subdomain1example")
		assert.Error(t, err)

	})
}
