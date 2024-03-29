// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
)

// StagedFileCreate is the builder for creating a StagedFile entity.
type StagedFileCreate struct {
	config
	mutation *StagedFileMutation
	hooks    []Hook
}

// SetPath sets the "path" field.
func (sfc *StagedFileCreate) SetPath(s string) *StagedFileCreate {
	sfc.mutation.SetPath(s)
	return sfc
}

// SetContent sets the "content" field.
func (sfc *StagedFileCreate) SetContent(s string) *StagedFileCreate {
	sfc.mutation.SetContent(s)
	return sfc
}

// Mutation returns the StagedFileMutation object of the builder.
func (sfc *StagedFileCreate) Mutation() *StagedFileMutation {
	return sfc.mutation
}

// Save creates the StagedFile in the database.
func (sfc *StagedFileCreate) Save(ctx context.Context) (*StagedFile, error) {
	return withHooks(ctx, sfc.sqlSave, sfc.mutation, sfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sfc *StagedFileCreate) SaveX(ctx context.Context) *StagedFile {
	v, err := sfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sfc *StagedFileCreate) Exec(ctx context.Context) error {
	_, err := sfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfc *StagedFileCreate) ExecX(ctx context.Context) {
	if err := sfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sfc *StagedFileCreate) check() error {
	if _, ok := sfc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "StagedFile.path"`)}
	}
	if v, ok := sfc.mutation.Path(); ok {
		if err := stagedfile.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "StagedFile.path": %w`, err)}
		}
	}
	if _, ok := sfc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "StagedFile.content"`)}
	}
	return nil
}

func (sfc *StagedFileCreate) sqlSave(ctx context.Context) (*StagedFile, error) {
	if err := sfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sfc.mutation.id = &_node.ID
	sfc.mutation.done = true
	return _node, nil
}

func (sfc *StagedFileCreate) createSpec() (*StagedFile, *sqlgraph.CreateSpec) {
	var (
		_node = &StagedFile{config: sfc.config}
		_spec = sqlgraph.NewCreateSpec(stagedfile.Table, sqlgraph.NewFieldSpec(stagedfile.FieldID, field.TypeInt))
	)
	if value, ok := sfc.mutation.Path(); ok {
		_spec.SetField(stagedfile.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := sfc.mutation.Content(); ok {
		_spec.SetField(stagedfile.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// StagedFileCreateBulk is the builder for creating many StagedFile entities in bulk.
type StagedFileCreateBulk struct {
	config
	builders []*StagedFileCreate
}

// Save creates the StagedFile entities in the database.
func (sfcb *StagedFileCreateBulk) Save(ctx context.Context) ([]*StagedFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sfcb.builders))
	nodes := make([]*StagedFile, len(sfcb.builders))
	mutators := make([]Mutator, len(sfcb.builders))
	for i := range sfcb.builders {
		func(i int, root context.Context) {
			builder := sfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StagedFileMutation)
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
					_, err = mutators[i+1].Mutate(root, sfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sfcb *StagedFileCreateBulk) SaveX(ctx context.Context) []*StagedFile {
	v, err := sfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sfcb *StagedFileCreateBulk) Exec(ctx context.Context) error {
	_, err := sfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfcb *StagedFileCreateBulk) ExecX(ctx context.Context) {
	if err := sfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
