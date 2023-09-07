// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/global/ent/logintoken"
)

// LoginToken is the model entity for the LoginToken schema.
type LoginToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SignInUserSlug holds the value of the "sign_in_user_slug" field.
	SignInUserSlug string `json:"sign_in_user_slug,omitempty"`
	// CliLoginToken holds the value of the "cli_login_token" field.
	CliLoginToken string `json:"cli_login_token,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LoginToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case logintoken.FieldID:
			values[i] = new(sql.NullInt64)
		case logintoken.FieldSignInUserSlug, logintoken.FieldCliLoginToken:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoginToken fields.
func (lt *LoginToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case logintoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lt.ID = int(value.Int64)
		case logintoken.FieldSignInUserSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign_in_user_slug", values[i])
			} else if value.Valid {
				lt.SignInUserSlug = value.String
			}
		case logintoken.FieldCliLoginToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cli_login_token", values[i])
			} else if value.Valid {
				lt.CliLoginToken = value.String
			}
		default:
			lt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LoginToken.
// This includes values selected through modifiers, order, etc.
func (lt *LoginToken) Value(name string) (ent.Value, error) {
	return lt.selectValues.Get(name)
}

// Update returns a builder for updating this LoginToken.
// Note that you need to call LoginToken.Unwrap() before calling this method if this LoginToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (lt *LoginToken) Update() *LoginTokenUpdateOne {
	return NewLoginTokenClient(lt.config).UpdateOne(lt)
}

// Unwrap unwraps the LoginToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lt *LoginToken) Unwrap() *LoginToken {
	_tx, ok := lt.config.driver.(*txDriver)
	if !ok {
		panic("ent: LoginToken is not a transactional entity")
	}
	lt.config.driver = _tx.drv
	return lt
}

// String implements the fmt.Stringer.
func (lt *LoginToken) String() string {
	var builder strings.Builder
	builder.WriteString("LoginToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lt.ID))
	builder.WriteString("sign_in_user_slug=")
	builder.WriteString(lt.SignInUserSlug)
	builder.WriteString(", ")
	builder.WriteString("cli_login_token=")
	builder.WriteString(lt.CliLoginToken)
	builder.WriteByte(')')
	return builder.String()
}

// LoginTokens is a parsable slice of LoginToken.
type LoginTokens []*LoginToken