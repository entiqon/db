# Identifier Types

> Part of [Entiqon](https://github.com/entiqon) / [Database](../../../) / [Token](../../) / [Types](../)

The `identifier` package classifies SQL expressions into broad syntactic
categories. This classification is purely **syntactic**, not semantic,
and is used internally by token resolvers and helpers.

---

## Purpose

* Provide a dependency-free enum (`identifier.Type`) to represent
  the form of a SQL expression.
* Enable higher-level tokens (`Field`, `Table`, …) and builders
  (`SelectBuilder`, …) to share consistent semantics without cycles.
* Support helpers such as `ClassifyExpression` to normalize input
  into one of the defined categories.

---

## Categories

The following categories are supported:

| Constant         | Description                                   | Example                      |
|------------------|-----------------------------------------------|------------------------------|
| `TypeInvalid`    | Could not classify                            | `""`                         |
| `TypeSubquery`   | Parenthesized SELECT                          | `(SELECT * FROM users)`      |
| `TypeComputed`   | Other parenthesized expression                | `(a + b)`                    |
| `TypeAggregate`  | Aggregate function                            | `SUM(qty)`, `COUNT(*)`       |
| `TypeFunction`   | Any other function or call                    | `JSON_EXTRACT(data, '$.id')` |
| `TypeLiteral`    | Quoted string or numeric constant             | `'abc'`, `"xyz"`, `42`       |
| `TypeExpression` | Plain table or column name (default fallback) | `users`, `id`                |
| `TypeWildcard`   | Wildcard symbol or qualified form             | `*`, `table.*`               |

---

## Public API

The following exported methods are available on `identifier.Type`:

### `func (k Type) String() string`

Returns the canonical name of the classification (e.g. `"Function"`, `"Literal"`).  
If the type is not recognized or invalid, returns `"Invalid"`.

---

### `func (k Type) Alias() string`

Returns the short alias (two-letter code) used when generating automatic
aliases for expressions.

| Type            | Alias |
|-----------------|-------|
| `TypeFunction`  | `fn`  |
| `TypeAggregate` | `ag`  |
| `TypeLiteral`   | `lt`  |
| `TypeExpression`| `ex`  |
| `TypeWildcard`  | `wc`  |
| `TypeComputed`  | `cp`  |
| `TypeSubquery`  | `sq`  |

---

### `func (k Type) IsValid() bool`

Reports whether the type is valid and registered in the internal registry.  
Invalid and unknown types return `false`.

---

### `func (k Type) IsWildcard() bool`

Reports whether the type represents a wildcard expression, such as `*` or
a qualified form like `table.*`.

---

### `func ParseType(value any) Type`

Attempts to coerce an arbitrary value into a `Type`.

Accepted inputs:
- `Type`: returned directly if valid
- `int`: cast to `Type` if valid
- `string`: matched case-insensitively against both canonical names and aliases

Returns `TypeInvalid` if no match is found.

---

## Example

```go
package main

import (
    "fmt"
    "github.com/entiqon/db/token/types/identifier"
)

func main() {
    var t identifier.Type

    // Identifier types are prefixed with Type for clarity.
    t = identifier.TypeFunction
    fmt.Println(t) // Function

    // Check validity
    fmt.Println(t.IsValid()) // true

    // Aliases
    fmt.Println(t.Alias()) // fn

    // Wildcard example
    w := identifier.TypeWildcard
    fmt.Println(w.IsWildcard()) // true

    // Parsing from string
    parsed := identifier.ParseType("ag")
    fmt.Println(parsed) // Aggregate
}
```

---

## License

Released under the [MIT License](../../../../LICENSE).  
Copyright © 2025 [Entiqon Contributors](https://entiqon.dev)
