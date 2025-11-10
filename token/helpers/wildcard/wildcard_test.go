package wildcard_test

import (
	"testing"

	"github.com/entiqon/db/token/helpers/wildcard"
)

func TestWildcard(t *testing.T) {
	t.Run("Methods", func(t *testing.T) {
		tests := []struct {
			Name     string
			Input    string
			WantBase string
			WantIs   bool
			WantErr  bool
			Valid    bool
		}{
			// ParseWildcard
			{Name: "BareWildcard", Input: "*", WantBase: "*", WantIs: true, WantErr: false, Valid: true},
			{Name: "QualifiedWildcard", Input: "users.*", WantBase: "users.*", WantIs: true, WantErr: false, Valid: true},
			{Name: "AliasedWildcard", Input: "* alias", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "AliasedWildcardWithKeyword", Input: "* AS alias", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "QualifiedAliasedWildcard", Input: "users.* AS alias", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "AliasKeyword", Input: "* AS", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "ExtraSpaces", Input: "  *  ", WantBase: "*", WantIs: true, WantErr: false, Valid: true},
			{Name: "NonWildcardExpression", Input: "qty * price", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "EmptyString", Input: "", WantBase: "", WantIs: false, WantErr: true, Valid: false},
			{Name: "EmptyValue", Input: "  ", WantBase: "", WantIs: false, WantErr: true, Valid: false},
		}

		t.Run("ParseWildcard", func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.Name, func(t *testing.T) {
					base, err := wildcard.ParseWildcard(tt.Input)
					if base != tt.WantBase {
						t.Errorf("ParseWildcard(%q) base = %q, want %q", tt.Input, base, tt.WantBase)
					}
					if (err != nil) != tt.WantErr {
						t.Errorf("ParseWildcard(%q) error = %v, wantErr %v", tt.Input, err, tt.WantErr)
					}
				})
			}
		})

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
			{Name: "MixedCaseAlias", Input: "Users.* AS Total", Valid: false},  // aliased → invalid
			{Name: "IrregularSpacing", Input: "users  .  *", Valid: false},     // malformed
			{Name: "NestedWildcardLike", Input: "(users.*)", Valid: false},     // not a wildcard
			{Name: "LowercaseAs", Input: "users.* as alias", Valid: false},     // aliased
			{Name: "SchemaQualified", Input: "public.users.*", Valid: true},    // qualified valid
			{Name: "InvalidSyntaxDoubleStar", Input: "users.**", Valid: false}, // bad syntax
			{Name: "InvalidSyntaxTrailing", Input: "users.*extra", Valid: false},
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
