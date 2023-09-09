package src

import (
	"fmt"
	"net/url"
	"strings"
)

type HostVal struct {
	source string
}

func (v *HostVal) String() string {
	if !v.isValid() {
		return ""
	}

	return v.source
}

func (v *HostVal) isValid() bool {
	prefix := "cspolicy.src.HostSource"

	parsedURL, err := url.Parse(v.source)
	if err != nil {
		format := "[%s] '%s' is invalid, skipping... (url.Parse error: '%v')\n"
		fmt.Printf(format, prefix, v.source, err)

		return false
	}

	var padScheme string
	if parsedURL.Scheme == "" {
		padScheme = "noscheme://"
	}

	_, err = url.ParseRequestURI(padScheme + v.source)
	if err != nil {
		preparedErr := strings.ReplaceAll(err.Error(), padScheme, "")

		format := "[%s] '%s' is invalid, skipping... (url.ParseRequestURI error: '%v')\n"
		fmt.Printf(format, prefix, v.source, preparedErr)

		return false
	}

	return err == nil
}

func Host(source string) *HostVal {
	return &HostVal{source}
}
