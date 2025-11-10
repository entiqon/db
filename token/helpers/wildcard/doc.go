// Package wildcard provides lexical utilities for detecting and validating
// SQL wildcard expressions such as "*" and "table.*".
//
// # Overview
//
// Wildcards are a special case of SQL expressions that denote column expansion
// at query rendering time. The helpers in this package perform *syntactic*
// checks only — they do not access schema metadata or perform column resolution.
//
// Core Behaviors
//
//   - A valid wildcard is either "*" or a qualified form like "users.*".
//   - Aliased wildcards such as "* AS alias" or "users.* alias" are rejected.
//   - Incomplete or malformed expressions like "* AS" are also rejected.
//   - Non-wildcard formulas (e.g., "qty * price") are classified as invalid.
//
// These helpers are typically used by query builders and token analyzers to
// distinguish wildcard selectors from ordinary identifiers or computed columns.
//
// Functions
//
//   - ParseWildcard(expr string) (bool, error)
//     Performs strict validation of a potential wildcard expression and returns
//     whether it is valid along with an explanatory error if not.
//
//   - IsWildcard(expr string) bool
//     Lightweight check that reports true only for syntactically valid, unaliased
//     wildcards.
//
//   - ValidateWildcard(expr string) error
//     Returns nil if the expression is valid or a descriptive error otherwise.
//
// Example
//
//	ok, err := wildcard.ParseWildcard("users.*")
//	if ok {
//	    fmt.Println("Valid wildcard")
//	} else {
//	    fmt.Println("Invalid:", err)
//	}
package wildcard
