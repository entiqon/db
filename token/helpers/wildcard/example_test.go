package wildcard_test

import (
	"fmt"

	"github.com/entiqon/db/token/helpers/wildcard"
)

// ExampleParseWildcard demonstrates how to validate and classify
// wildcard expressions using the ParseWildcard function.
func ExampleParseWildcard() {
	cases := []string{
		"*",
		"users.*",
		"* AS alias",
		"users.* alias",
		"qty * price",
		"users.**",
	}

	for _, expr := range cases {
		ok, err := wildcard.ParseWildcard(expr)
		if ok {
			fmt.Printf("%-16q → valid\n", expr)
		} else {
			fmt.Printf("%-16q → invalid: %v\n", expr, err)
		}
	}

	// Output:
	// "*"              → valid
	// "users.*"        → valid
	// "* AS alias"     → invalid: wildcard cannot be aliased: "* AS alias"
	// "users.* alias"  → invalid: wildcard cannot be aliased: "users.* alias"
	// "qty * price"    → invalid: invalid wildcard syntax: "qty * price"
	// "users.**"       → invalid: invalid wildcard syntax: "users.**"
}

// ExampleIsWildcard shows how to perform lightweight wildcard detection
// without error details.
func ExampleIsWildcard() {
	fmt.Println(wildcard.IsWildcard("*"))
	fmt.Println(wildcard.IsWildcard("users.*"))
	fmt.Println(wildcard.IsWildcard("* AS alias"))
	fmt.Println(wildcard.IsWildcard("qty * price"))

	// Output:
	// true
	// true
	// false
	// false
}

// ExampleValidateWildcard demonstrates using ValidateWildcard to enforce
// wildcard correctness in SQL-like expressions.
func ExampleValidateWildcard() {
	if err := wildcard.ValidateWildcard("users.*"); err == nil {
		fmt.Println("Valid wildcard")
	}

	if err := wildcard.ValidateWildcard("* AS alias"); err != nil {
		fmt.Println("Invalid:", err)
	}

	// Output:
	// Valid wildcard
	// Invalid: wildcard cannot be aliased: "* AS alias"
}
