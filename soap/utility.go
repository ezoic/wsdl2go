package soap

import "strings"

// TrimNamespace removes the namespace from a SOAP typename
func TrimNamespace(s string) string {
	n := strings.SplitN(s, ":", 2)
	if len(n) == 2 {
		return n[1]
	}
	return s
}
