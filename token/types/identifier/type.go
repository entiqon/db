package identifier

import "strings"

// Type represents the syntactic classification of a SQL expression.
//
// Resolution categories:
//   - Invalid:    could not classify
//   - Subquery:   "(SELECT ...)"
//   - Computed:   other parenthesized expressions, e.g. "(a+b)"
//   - Aggregate:  SUM, COUNT, MAX, MIN, AVG
//   - Function:   other calls with parentheses, e.g. JSON_EXTRACT(data)
//   - Literal:    quoted string or numeric constant
//   - Expression: plain name (default fallback)
//   - Wildcard:   "*" or qualified form "table.*"
type Type int

const (
	TypeUnknown Type = iota
	TypeInvalid
	TypeSubquery
	TypeComputed
	TypeAggregate
	TypeFunction
	TypeLiteral
	TypeExpression
	TypeWildcard
)

// typeMeta holds metadata for a Type classification.
//
// Each SQL identifier type has:
//   - name:  the canonical human-readable label, used by String()
//   - alias: a short two-letter code, used for automatic alias generation
type typeMeta struct {
	name  string
	alias string
}

// registry centralizes all recognized SQL identifier types.
//
// It defines the authoritative mapping between a Type constant
// and its metadata (name + alias). All public methods such as
// String(), Alias(), IsValid(), and ParseFrom derive their values
// from this registry, ensuring consistency across the package.
//
// Adding a new Type only requires updating this registry.
var registry = map[Type]typeMeta{
	TypeExpression: {"Expression", "ex"},
	TypeWildcard:   {"Wildcard", "wc"},
	TypeLiteral:    {"Literal", "lt"},
	TypeFunction:   {"Function", "fn"},
	TypeAggregate:  {"Aggregate", "ag"},
	TypeComputed:   {"Computed", "cp"},
	TypeSubquery:   {"Subquery", "sq"},
}

// Alias returns the short two-letter code used when generating
// automatic aliases for this expression kind.
//
// If the Type is not registered (including Invalid),
// it returns an empty string.
func (k Type) Alias() string {
	if meta, ok := registry[k]; ok {
		return meta.alias
	}
	return ""
}

// IsWildcard reports whether the Type corresponds to a wildcard identifier.
//
// Wildcards include both unqualified "*" and qualified forms such as "table.*".
func (k Type) IsWildcard() bool {
	return k == TypeWildcard
}

// IsValid reports whether the Type is registered and not Unknown or Invalid.
func (k Type) IsValid() bool {
	if k == TypeInvalid || k == TypeUnknown {
		return false
	}
	_, ok := registry[k]
	return ok
}

// String returns the canonical label for the Type.
// If the Type is not registered (including Invalid),
// it returns an empty string.
func (k Type) String() string {
	if meta, ok := registry[k]; ok {
		return meta.name
	}
	return "Invalid"
}

// ParseFrom attempts to coerce an arbitrary value into a Type.
//
// Accepted inputs:
//   - Type: returned directly if valid
//   - int: cast to Type if valid
//   - string: matched case-insensitively against both the
//     canonical String() label and the Alias()
//
// Otherwise, returns Invalid.
func ParseFrom(value any) Type {
	switch v := value.(type) {
	case Type:
		if v.IsValid() {
			return v
		}
	case int:
		t := Type(v)
		if t.IsValid() {
			return t
		}
	case string:
		s := strings.TrimSpace(strings.ToLower(v))
		for t, meta := range registry {
			if s == strings.ToLower(meta.name) || s == strings.ToLower(meta.alias) {
				if t.IsValid() {
					return t
				}
			}
		}
	}
	return TypeInvalid
}
