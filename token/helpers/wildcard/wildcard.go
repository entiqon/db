package wildcard

import (
	"fmt"
	"strings"
)

// ParseWildcard validates whether the given expression represents a valid
// SQL wildcard pattern ("*" or "table.*").
//
// It performs strict lexical validation only—no schema resolution.
//
// It returns:
//
//   - ok = true, err = nil      → valid wildcard (unaliased)
//   - ok = false, err != nil    → invalid, malformed, or aliased
//
// Examples:
//
//	ParseWildcard("*")            → true, nil
//	ParseWildcard("users.*")      → true, nil
//	ParseWildcard("* AS alias")   → false, "wildcard cannot be aliased: * AS alias"
//	ParseWildcard("users.* AS a") → false, "wildcard cannot be aliased: users.* AS a"
//	ParseWildcard("* AS")         → false, "malformed expression: * AS"
//	ParseWildcard("qty * price")  → false, "invalid wildcard syntax: qty * price"
func ParseWildcard(expr string) (string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return "", fmt.Errorf("empty expression")
	}

	tokens := strings.Fields(expr)
	base := tokens[0]

	// Explicit alias form: "* AS alias" or "table.* AS alias"
	if len(tokens) >= 2 && strings.EqualFold(tokens[1], "AS") {
		if len(tokens) < 3 {
			return "", fmt.Errorf("malformed expression: %q", expr)
		}
		return "", fmt.Errorf("wildcard cannot be aliased: %q", expr)
	}

	// Implicit alias form: "* alias" or "table.* alias"
	if len(tokens) == 2 {
		return "", fmt.Errorf("wildcard cannot be aliased: %q", expr)
	}

	// Valid wildcard forms
	if base == "*" || strings.HasSuffix(base, ".*") {
		return base, nil
	}

	return "", fmt.Errorf("invalid wildcard syntax: %q", expr)
}

// IsWildcard reports whether expr is a syntactically valid unaliased wildcard.
//
// It simply wraps ParseWildcard and ignores the error detail.
// Returns true only if ParseWildcard returns ok == true.
func IsWildcard(expr string) bool {
	_, err := ParseWildcard(expr)
	return err == nil
}

// ValidateWildcard enforces correctness of wildcard usage.
//
// It returns nil if expr is a valid unaliased wildcard.
// Otherwise, it returns an error describing why it is invalid.
//
// Typical use: validation of SELECT field lists or column selectors
// before query rendering.
func ValidateWildcard(expr string) error {
	_, err := ParseWildcard(expr)
	if err != nil {
		return err
	}
	return nil
}
