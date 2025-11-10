package wildcard_test

import (
	"testing"

	"github.com/entiqon/db/token/helpers/wildcard"
)

func TestWildcard(t *testing.T) {
	t.Run("Methods", func(t *testing.T) {
		tests := []struct {
			Name      string
			Input     string
			WantBase  string
			WantAlias string
			WantOK    bool
			WantIs    bool
			WantErr   bool
			Valid     bool
		}{
			// Constructor (ParseWildcard)
			{Name: "BareWildcard", Input: "*", WantBase: "*", WantAlias: "", WantOK: true, WantIs: true, WantErr: false, Valid: true},
			{Name: "QualifiedWildcard", Input: "users.*", WantBase: "users.*", WantAlias: "", WantOK: true, WantIs: true, WantErr: false, Valid: true},
			{Name: "AliasedWildcard", Input: "* AS alias", WantBase: "*", WantAlias: "alias", WantOK: false, WantIs: false, WantErr: true, Valid: false},
			{Name: "QualifiedAliasedWildcard", Input: "users.* AS alias", WantBase: "users.*", WantAlias: "alias", WantOK: false, WantIs: false, WantErr: true, Valid: false},
			{Name: "AliasKeyword", Input: "* AS", WantBase: "", WantAlias: "", WantOK: false, WantIs: false, WantErr: true, Valid: false},
			{Name: "ExtraSpaces", Input: "  *  ", WantBase: "*", WantAlias: "", WantOK: true, WantIs: true, WantErr: false, Valid: true},
			{Name: "NonWildcardExpression", Input: "qty * price", WantBase: "", WantAlias: "", WantOK: false, WantIs: false, WantErr: true, Valid: false},
			{Name: "EmptyString", Input: "", WantBase: "", WantAlias: "", WantOK: false, WantIs: false, WantErr: true, Valid: false},
		}

		t.Run("IsWildcard", func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.Name, func(t *testing.T) {
					got := wildcard.IsWildcard(tt.Input)
					if got != tt.WantIs {
						t.Errorf("IsWildcard(%q) = %v, want %v", tt.Input, got, tt.WantIs)
					}
				})
			}
		})

		t.Run("ParseWildcard", func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.Name, func(t *testing.T) {
					base, alias, ok := wildcard.ParseWildcard(tt.Input)
					if ok != tt.WantOK || base != tt.WantBase || alias != tt.WantAlias {
						t.Errorf("ParseWildcard(%q) = (%q, %q, %v), want (%q, %q, %v)",
							tt.Input, base, alias, ok, tt.WantBase, tt.WantAlias, tt.WantOK)
					}
				})
			}
		})

		t.Run("ValidateWildcard", func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.Name, func(t *testing.T) {
					err := wildcard.ValidateWildcard(tt.Input)
					if (err != nil) != tt.WantErr {
						t.Errorf("ValidateWildcard(%q) error = %v, wantErr %v", tt.Input, err, tt.WantErr)
					}
				})
			}
		})
	})

	t.Run("EdgeCases", func(t *testing.T) {
		extras := []struct {
			Name  string
			Input string
			Valid bool
		}{
			{Name: "MixedCaseAlias", Input: "Users.* AS Total", Valid: false}, // aliased → invalid
			{Name: "IrregularSpacing", Input: "users  .  *", Valid: false},    // malformed
			{Name: "NestedWildcardLike", Input: "(users.*)", Valid: false},    // not a wildcard
			{Name: "LowercaseAs", Input: "users.* as alias", Valid: false},    // aliased
			{Name: "SchemaQualified", Input: "public.users.*", Valid: true},   // qualified valid
		}

		for _, tt := range extras {
			t.Run(tt.Name, func(t *testing.T) {
				err := wildcard.ValidateWildcard(tt.Input)
				if (err == nil) != tt.Valid {
					t.Errorf("ValidateWildcard(%q) validity = %v, want %v", tt.Input, err == nil, tt.Valid)
				}
			})
		}
	})
}
