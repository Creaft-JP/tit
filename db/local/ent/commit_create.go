// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/commit"
	"github.com/Creaft-JP/tit/db/local/ent/committedfile"
)

// CommitCreate is the builder for creating a Commit entity.
type CommitCreate struct {
	config
	mutation *CommitMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (cc *CommitCreate) SetNumber(i int) *CommitCreate {
	cc.mutation.SetNumber(i)
	return cc
}

// SetMessage sets the "message" field.
func (cc *CommitCreate) SetMessage(s string) *CommitCreate {
	cc.mutation.SetMessage(s)
	return cc
}

// AddFileIDs adds the "files" edge to the CommittedFile entity by IDs.
func (cc *CommitCreate) AddFileIDs(ids ...int) *CommitCreate {
	cc.mutation.AddFileIDs(ids...)
	return cc
}

// AddFiles adds the "files" edges to the CommittedFile entity.
func (cc *CommitCreate) AddFiles(c ...*CommittedFile) *CommitCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddFileIDs(ids...)
}

// Mutation returns the CommitMutation object of the builder.
func (cc *CommitCreate) Mutation() *CommitMutation {
	return cc.mutation
}

// Save creates the Commit in the database.
func (cc *CommitCreate) Save(ctx context.Context) (*Commit, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommitCreate) SaveX(ctx context.Context) *Commit {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommitCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommitCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommitCreate) check() error {
	if _, ok := cc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Commit.number"`)}
	}
	if v, ok := cc.mutation.Number(); ok {
		if err := commit.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Commit.number": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "Commit.message"`)}
	}
	if v, ok := cc.mutation.Message(); ok {
		if err := commit.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Commit.message": %w`, err)}
		}
	}
	return nil
}

func (cc *CommitCreate) sqlSave(ctx context.Context) (*Commit, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommitCreate) createSpec() (*Commit, *sqlgraph.CreateSpec) {
	var (
		_node = &Commit{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(commit.Table, sqlgraph.NewFieldSpec(commit.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Number(); ok {
		_spec.SetField(commit.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := cc.mutation.Message(); ok {
		_spec.SetField(commit.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if nodes := cc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commit.FilesTable,
			Columns: []string{commit.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(committedfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommitCreateBulk is the builder for creating many Commit entities in bulk.
type CommitCreateBulk struct {
	config
	builders []*CommitCreate
}

// Save creates the Commit entities in the database.
func (ccb *CommitCreateBulk) Save(ctx context.Context) ([]*Commit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Commit, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommitMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommitCreateBulk) SaveX(ctx context.Context) []*Commit {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommitCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommitCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
