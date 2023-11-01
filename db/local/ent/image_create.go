// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/image"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
	"github.com/google/uuid"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
}

// SetExtension sets the "extension" field.
func (ic *ImageCreate) SetExtension(s string) *ImageCreate {
	ic.mutation.SetExtension(s)
	return ic
}

// SetContents sets the "contents" field.
func (ic *ImageCreate) SetContents(b []byte) *ImageCreate {
	ic.mutation.SetContents(b)
	return ic
}

// SetNumber sets the "number" field.
func (ic *ImageCreate) SetNumber(i int) *ImageCreate {
	ic.mutation.SetNumber(i)
	return ic
}

// SetDescription sets the "description" field.
func (ic *ImageCreate) SetDescription(s string) *ImageCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetID sets the "id" field.
func (ic *ImageCreate) SetID(u uuid.UUID) *ImageCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *ImageCreate) SetNillableID(u *uuid.UUID) *ImageCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// AddCommitIDs adds the "commit" edge to the TitCommit entity by IDs.
func (ic *ImageCreate) AddCommitIDs(ids ...int) *ImageCreate {
	ic.mutation.AddCommitIDs(ids...)
	return ic
}

// AddCommit adds the "commit" edges to the TitCommit entity.
func (ic *ImageCreate) AddCommit(t ...*TitCommit) *ImageCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ic.AddCommitIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImageCreate) defaults() {
	if _, ok := ic.mutation.ID(); !ok {
		v := image.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.Extension(); !ok {
		return &ValidationError{Name: "extension", err: errors.New(`ent: missing required field "Image.extension"`)}
	}
	if v, ok := ic.mutation.Extension(); ok {
		if err := image.ExtensionValidator(v); err != nil {
			return &ValidationError{Name: "extension", err: fmt.Errorf(`ent: validator failed for field "Image.extension": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`ent: missing required field "Image.contents"`)}
	}
	if _, ok := ic.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Image.number"`)}
	}
	if _, ok := ic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Image.description"`)}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(image.Table, sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.Extension(); ok {
		_spec.SetField(image.FieldExtension, field.TypeString, value)
		_node.Extension = value
	}
	if value, ok := ic.mutation.Contents(); ok {
		_spec.SetField(image.FieldContents, field.TypeBytes, value)
		_node.Contents = value
	}
	if value, ok := ic.mutation.Number(); ok {
		_spec.SetField(image.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(image.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := ic.mutation.CommitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   image.CommitTable,
			Columns: image.CommitPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(titcommit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	err      error
	builders []*ImageCreate
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
