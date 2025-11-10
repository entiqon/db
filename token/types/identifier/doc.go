// Package identifier provides the classification of SQL expressions
// into broad syntactic categories such as subqueries, functions,
// aggregates, literals, and plain identifiers.
//
// # Overview
//
// The identifier.Type enum is a low-level building block used by
// SQL tokens (Field, Table, …) and higher-level builders
// (SelectBuilder, …). It allows consistent parsing and validation
// of input expressions without introducing cyclic dependencies.
//
// Classification is purely syntactic, not semantic. For example,
// SUM(qty) is classified as an Aggregate even if it appears in an
// invalid position in the query.
//
// # Categories
//
//   - TypeInvalid:    could not classify
//   - TypeSubquery:   "(SELECT ...)"
//   - TypeComputed:   other parenthesized expressions, e.g. "(a+b)"
//   - TypeAggregate:  SUM, COUNT, MAX, MIN, AVG
//   - TypeFunction:   other calls with parentheses, e.g. JSON_EXTRACT(data)
//   - TypeLiteral:    quoted string or numeric constant
//   - TypeExpression: plain table or column name (default fallback)
//   - TypeWildcard:   wildcard symbol or qualified form (e.g. table.*)
//
// # Philosophy
//
//   - Never panic: always return a Type, with TypeInvalid or TypeUnknown
//     as safe fallbacks.
//   - Auditability: preserve the original classification for
//     debugging and logs.
//   - Strict enforcement: higher-level resolvers must reject inputs
//     that do not classify correctly.
//
// Example usage is provided in example_test.go and the package README.
package identifier
