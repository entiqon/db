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

		if got := identifier.Wildcard.Alias(); got != "wc" {
			t.Errorf("Wildcard.Alias() = %v, want %v", got, "wc")
		}
	})

	t.Run("IsValid", func(t *testing.T) {
		if identifier.ParseFrom("").IsValid() {
			t.Errorf("expect to be invalid")
		}

		if !identifier.Wildcard.IsValid() {
			t.Errorf("expect to be valid")
		}
	})

	t.Run("String", func(t *testing.T) {
		if got := identifier.ParseFrom("").String(); got != identifier.Invalid.String() {
			t.Errorf("Identifier.ParseFrom()=%q, want empty", got)
		}

		if got := identifier.ParseFrom(7).String(); got != identifier.Wildcard.String() {
			t.Errorf("Identifier.ParseFrom()=%q, want empty", got)
		}

		if got := identifier.Wildcard.Alias(); got != "wc" {
			t.Errorf("Wildcard.Alias() = %v, want %v", got, "wc")
		}
	})

	t.Run("ParseFrom", func(t *testing.T) {
		if identifier.ParseFrom("") != identifier.Invalid {
			t.Errorf("expect to be invalid")
		}

		if got := identifier.ParseFrom([]string{"bad"}); got != identifier.Invalid {
			t.Errorf("ParseFrom([]string) = %v, want Invalid", got)
		}

		if identifier.ParseFrom("wc") != identifier.Wildcard {
			t.Errorf("expect to be wc")
		}

		if identifier.ParseFrom("wildcard") != identifier.Wildcard {
			t.Errorf("expect to be wildcard")
		}

		if identifier.ParseFrom("WILDCARD") != identifier.Wildcard {
			t.Errorf("expect to be wildcard")
		}

		if got := identifier.ParseFrom(identifier.Computed); got != identifier.Computed {
			t.Errorf("ParseFrom(Type Computed) = %v, want Computed", got)
		}
	})
}
