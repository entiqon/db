# 📘 Wildcard Helper

> Part of [Entiqon](../../../) / [Database](../../) / [Token](../)
> 
> **Purpose:** Lightweight lexical validator for SQL-style wildcard expressions such as `*` or `table.*`.

---

## Overview

The **wildcard** package provides syntactic validation and classification of SQL wildcard expressions.  
It’s designed for use in token analyzers, expression resolvers, and query builders that must distinguish  
between ordinary identifiers and column-expanding wildcard selectors.

The logic is purely lexical — it does **not** access schema metadata or verify column existence.

---

## ✨ Features

- ✅ Detects valid wildcard expressions (`*`, `users.*`, `public.table.*`)
- 🚫 Rejects aliased or malformed variants (`* AS alias`, `table.* alias`, `* AS`)
- 🧩 Distinguishes true wildcards from arithmetic or non-wildcard expressions (`qty * price`)
- ⚙️ Produces consistent, descriptive errors for invalid syntax or alias misuse
- 🧪 Fully covered by unit tests (`wildcard_test.go`)

---

## ⚙️ API

### `ParseWildcard(expr string) (bool, error)`

Performs strict validation of a potential wildcard expression.

| Input           | ok | Error                        |
|-----------------|----|------------------------------|
| `"*"`           | ✅  | —                            |
| `"users.*"`     | ✅  | —                            |
| `"* AS alias"`  | ❌  | `wildcard cannot be aliased` |
| `"* AS"`        | ❌  | `malformed expression`       |
| `"qty * price"` | ❌  | `invalid wildcard syntax`    |

---

### `IsWildcard(expr string) bool`

Returns `true` only for syntactically valid, unaliased wildcard expressions.

Example:
```go
if wildcard.IsWildcard("orders.*") {
    fmt.Println("Wildcard selector detected")
}
```

---

### `ValidateWildcard(expr string) error`

Convenience wrapper around `ParseWildcard`.  
Returns `nil` if valid, or a detailed error explaining why it isn’t.

Example:
```go
if err := wildcard.ValidateWildcard("* AS alias"); err != nil {
    log.Printf("invalid wildcard: %v", err)
}
```

---

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Sample output:
```
ok  	github.com/entiqon/db/token/helpers/wildcard	0.005s
```

---

## 🔗 Related Packages

- [`identifier`](../identifier) — identifier parsing and validation.
- [`expression`](../expression) — SQL expression resolution and classification.
- [`table`](../table) — source and alias handling.

---

## 🧠 Design Notes

- The package avoids regex parsing for simplicity and speed.
- All checks are whitespace- and case-insensitive for SQL keywords (`AS`, `as`, etc.).
- Error messages are short and deterministic, suitable for internal validation or user feedback.

---

## 📄 License

MIT © [ENTIQON Labs](https://entiqon.dev)
