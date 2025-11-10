package identifier_test

import (
	"fmt"

	"github.com/entiqon/db/token/types/identifier"
)

func ExampleType_Alias() {
	t := identifier.TypeWildcard
	fmt.Println(t.Alias())

	// Output:
	// wc
}

func ExampleType_IsValid() {
	t := identifier.TypeInvalid
	fmt.Println(t.IsValid())

	t = identifier.TypeUnknown
	fmt.Println(t.IsValid())

	t = identifier.TypeSubquery
	fmt.Println(t.IsValid())

	// Output:
	// false
	// false
	// true
}

func ExampleType_IsWildcard() {
	var t identifier.Type
	t = identifier.TypeFunction
	fmt.Println(t.IsWildcard())

	t = identifier.TypeWildcard
	fmt.Println(t.IsWildcard())

	// Output:
	// false
	// true
}

func ExampleType_String() {
	t := identifier.TypeFunction
	fmt.Println(t.String())

	// Output:
	// Function
}

func ExampleType_parseFrom() {
	t := identifier.ParseFrom(123)
	fmt.Println(t.String())

	t = identifier.ParseFrom(7)
	fmt.Println(t.String())

	// Output:
	// Invalid
	// Expression
}
