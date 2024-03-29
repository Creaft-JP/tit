// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/remote"
)

// RemoteCreate is the builder for creating a Remote entity.
type RemoteCreate struct {
	config
	mutation *RemoteMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RemoteCreate) SetName(s string) *RemoteCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetURL sets the "url" field.
func (rc *RemoteCreate) SetURL(s string) *RemoteCreate {
	rc.mutation.SetURL(s)
	return rc
}

// Mutation returns the RemoteMutation object of the builder.
func (rc *RemoteCreate) Mutation() *RemoteMutation {
	return rc.mutation
}

// Save creates the Remote in the database.
func (rc *RemoteCreate) Save(ctx context.Context) (*Remote, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RemoteCreate) SaveX(ctx context.Context) *Remote {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RemoteCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RemoteCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RemoteCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Remote.name"`)}
	}
	if _, ok := rc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Remote.url"`)}
	}
	if v, ok := rc.mutation.URL(); ok {
		if err := remote.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Remote.url": %w`, err)}
		}
	}
	return nil
}

func (rc *RemoteCreate) sqlSave(ctx context.Context) (*Remote, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RemoteCreate) createSpec() (*Remote, *sqlgraph.CreateSpec) {
	var (
		_node = &Remote{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(remote.Table, sqlgraph.NewFieldSpec(remote.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(remote.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.URL(); ok {
		_spec.SetField(remote.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	return _node, _spec
}

// RemoteCreateBulk is the builder for creating many Remote entities in bulk.
type RemoteCreateBulk struct {
	config
	builders []*RemoteCreate
}

// Save creates the Remote entities in the database.
func (rcb *RemoteCreateBulk) Save(ctx context.Context) ([]*Remote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Remote, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RemoteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RemoteCreateBulk) SaveX(ctx context.Context) []*Remote {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RemoteCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RemoteCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
