package identifier_test

import (
	"testing"

	"github.com/entiqon/db/token/types/identifier"
)

func TestType(t *testing.T) {
	t.Run("Alias", func(t *testing.T) {
		if got := identifier.ParseFrom("").Alias(); got != "" {
			t.Errorf("Identifier.ParseFrom()=%q, want empty", got)
		}

		if got := identifier.TypeWildcard.Alias(); got != "wc" {
			t.Errorf("Wildcard.Alias() = %v, want %v", got, "wc")
		}
	})

	t.Run("IsWildcard", func(t *testing.T) {
		if got := identifier.ParseFrom(8).IsWildcard(); !got {
			t.Errorf("IsWildcard() = %v, want %v", got, true)
		}
		if got := identifier.ParseFrom(9).IsWildcard(); got {
			t.Errorf("IsWildcard() = %v, want %v", got, false)
		}
	})

	t.Run("IsValid", func(t *testing.T) {
		if identifier.ParseFrom("").IsValid() {
			t.Errorf("expect to be invalid")
		}

		if !identifier.TypeWildcard.IsValid() {
			t.Errorf("expect to be valid")
		}
	})

	t.Run("String", func(t *testing.T) {
		if got := identifier.ParseFrom("").String(); got != identifier.TypeInvalid.String() {
			t.Errorf("Identifier.ParseFrom()=%q, want empty", got)
		}

		if got := identifier.ParseFrom(8).String(); got != identifier.TypeWildcard.String() {
			t.Errorf("Identifier.ParseFrom()=%q, want empty", got)
		}

		if got := identifier.TypeWildcard.Alias(); got != "wc" {
			t.Errorf("Wildcard.Alias() = %v, want %v", got, "wc")
		}
	})

	t.Run("ParseFrom", func(t *testing.T) {
		if identifier.ParseFrom("") != identifier.TypeInvalid {
			t.Errorf("expect to be invalid")
		}

		if got := identifier.ParseFrom([]string{"bad"}); got != identifier.TypeInvalid {
			t.Errorf("ParseFrom([]string) = %v, want Invalid", got)
		}

		if identifier.ParseFrom("wc") != identifier.TypeWildcard {
			t.Errorf("expect to be wc")
		}

		if identifier.ParseFrom("wildcard") != identifier.TypeWildcard {
			t.Errorf("expect to be wildcard")
		}

		if identifier.ParseFrom("WILDCARD") != identifier.TypeWildcard {
			t.Errorf("expect to be wildcard")
		}

		if got := identifier.ParseFrom(identifier.TypeComputed); got != identifier.TypeComputed {
			t.Errorf("ParseFrom(Type Computed) = %v, want Computed", got)
		}
	})
}
