// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/global/ent/logintoken"
	"github.com/Creaft-JP/tit/db/global/ent/predicate"
)

// LoginTokenUpdate is the builder for updating LoginToken entities.
type LoginTokenUpdate struct {
	config
	hooks    []Hook
	mutation *LoginTokenMutation
}

// Where appends a list predicates to the LoginTokenUpdate builder.
func (ltu *LoginTokenUpdate) Where(ps ...predicate.LoginToken) *LoginTokenUpdate {
	ltu.mutation.Where(ps...)
	return ltu
}

// SetSignInUserSlug sets the "sign_in_user_slug" field.
func (ltu *LoginTokenUpdate) SetSignInUserSlug(s string) *LoginTokenUpdate {
	ltu.mutation.SetSignInUserSlug(s)
	return ltu
}

// SetCliLoginToken sets the "cli_login_token" field.
func (ltu *LoginTokenUpdate) SetCliLoginToken(s string) *LoginTokenUpdate {
	ltu.mutation.SetCliLoginToken(s)
	return ltu
}

// Mutation returns the LoginTokenMutation object of the builder.
func (ltu *LoginTokenUpdate) Mutation() *LoginTokenMutation {
	return ltu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ltu *LoginTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ltu.sqlSave, ltu.mutation, ltu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltu *LoginTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := ltu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ltu *LoginTokenUpdate) Exec(ctx context.Context) error {
	_, err := ltu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltu *LoginTokenUpdate) ExecX(ctx context.Context) {
	if err := ltu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltu *LoginTokenUpdate) check() error {
	if v, ok := ltu.mutation.SignInUserSlug(); ok {
		if err := logintoken.SignInUserSlugValidator(v); err != nil {
			return &ValidationError{Name: "sign_in_user_slug", err: fmt.Errorf(`ent: validator failed for field "LoginToken.sign_in_user_slug": %w`, err)}
		}
	}
	if v, ok := ltu.mutation.CliLoginToken(); ok {
		if err := logintoken.CliLoginTokenValidator(v); err != nil {
			return &ValidationError{Name: "cli_login_token", err: fmt.Errorf(`ent: validator failed for field "LoginToken.cli_login_token": %w`, err)}
		}
	}
	return nil
}

func (ltu *LoginTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ltu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(logintoken.Table, logintoken.Columns, sqlgraph.NewFieldSpec(logintoken.FieldID, field.TypeInt))
	if ps := ltu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltu.mutation.SignInUserSlug(); ok {
		_spec.SetField(logintoken.FieldSignInUserSlug, field.TypeString, value)
	}
	if value, ok := ltu.mutation.CliLoginToken(); ok {
		_spec.SetField(logintoken.FieldCliLoginToken, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ltu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logintoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ltu.mutation.done = true
	return n, nil
}

// LoginTokenUpdateOne is the builder for updating a single LoginToken entity.
type LoginTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginTokenMutation
}

// SetSignInUserSlug sets the "sign_in_user_slug" field.
func (ltuo *LoginTokenUpdateOne) SetSignInUserSlug(s string) *LoginTokenUpdateOne {
	ltuo.mutation.SetSignInUserSlug(s)
	return ltuo
}

// SetCliLoginToken sets the "cli_login_token" field.
func (ltuo *LoginTokenUpdateOne) SetCliLoginToken(s string) *LoginTokenUpdateOne {
	ltuo.mutation.SetCliLoginToken(s)
	return ltuo
}

// Mutation returns the LoginTokenMutation object of the builder.
func (ltuo *LoginTokenUpdateOne) Mutation() *LoginTokenMutation {
	return ltuo.mutation
}

// Where appends a list predicates to the LoginTokenUpdate builder.
func (ltuo *LoginTokenUpdateOne) Where(ps ...predicate.LoginToken) *LoginTokenUpdateOne {
	ltuo.mutation.Where(ps...)
	return ltuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ltuo *LoginTokenUpdateOne) Select(field string, fields ...string) *LoginTokenUpdateOne {
	ltuo.fields = append([]string{field}, fields...)
	return ltuo
}

// Save executes the query and returns the updated LoginToken entity.
func (ltuo *LoginTokenUpdateOne) Save(ctx context.Context) (*LoginToken, error) {
	return withHooks(ctx, ltuo.sqlSave, ltuo.mutation, ltuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ltuo *LoginTokenUpdateOne) SaveX(ctx context.Context) *LoginToken {
	node, err := ltuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ltuo *LoginTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := ltuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltuo *LoginTokenUpdateOne) ExecX(ctx context.Context) {
	if err := ltuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltuo *LoginTokenUpdateOne) check() error {
	if v, ok := ltuo.mutation.SignInUserSlug(); ok {
		if err := logintoken.SignInUserSlugValidator(v); err != nil {
			return &ValidationError{Name: "sign_in_user_slug", err: fmt.Errorf(`ent: validator failed for field "LoginToken.sign_in_user_slug": %w`, err)}
		}
	}
	if v, ok := ltuo.mutation.CliLoginToken(); ok {
		if err := logintoken.CliLoginTokenValidator(v); err != nil {
			return &ValidationError{Name: "cli_login_token", err: fmt.Errorf(`ent: validator failed for field "LoginToken.cli_login_token": %w`, err)}
		}
	}
	return nil
}

func (ltuo *LoginTokenUpdateOne) sqlSave(ctx context.Context) (_node *LoginToken, err error) {
	if err := ltuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(logintoken.Table, logintoken.Columns, sqlgraph.NewFieldSpec(logintoken.FieldID, field.TypeInt))
	id, ok := ltuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LoginToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ltuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logintoken.FieldID)
		for _, f := range fields {
			if !logintoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logintoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ltuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ltuo.mutation.SignInUserSlug(); ok {
		_spec.SetField(logintoken.FieldSignInUserSlug, field.TypeString, value)
	}
	if value, ok := ltuo.mutation.CliLoginToken(); ok {
		_spec.SetField(logintoken.FieldCliLoginToken, field.TypeString, value)
	}
	_node = &LoginToken{config: ltuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ltuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logintoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ltuo.mutation.done = true
	return _node, nil
}