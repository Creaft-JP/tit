// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
)

// StagedFileQuery is the builder for querying StagedFile entities.
type StagedFileQuery struct {
	config
	ctx        *QueryContext
	order      []stagedfile.OrderOption
	inters     []Interceptor
	predicates []predicate.StagedFile
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StagedFileQuery builder.
func (sfq *StagedFileQuery) Where(ps ...predicate.StagedFile) *StagedFileQuery {
	sfq.predicates = append(sfq.predicates, ps...)
	return sfq
}

// Limit the number of records to be returned by this query.
func (sfq *StagedFileQuery) Limit(limit int) *StagedFileQuery {
	sfq.ctx.Limit = &limit
	return sfq
}

// Offset to start from.
func (sfq *StagedFileQuery) Offset(offset int) *StagedFileQuery {
	sfq.ctx.Offset = &offset
	return sfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sfq *StagedFileQuery) Unique(unique bool) *StagedFileQuery {
	sfq.ctx.Unique = &unique
	return sfq
}

// Order specifies how the records should be ordered.
func (sfq *StagedFileQuery) Order(o ...stagedfile.OrderOption) *StagedFileQuery {
	sfq.order = append(sfq.order, o...)
	return sfq
}

// First returns the first StagedFile entity from the query.
// Returns a *NotFoundError when no StagedFile was found.
func (sfq *StagedFileQuery) First(ctx context.Context) (*StagedFile, error) {
	nodes, err := sfq.Limit(1).All(setContextOp(ctx, sfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{stagedfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sfq *StagedFileQuery) FirstX(ctx context.Context) *StagedFile {
	node, err := sfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StagedFile ID from the query.
// Returns a *NotFoundError when no StagedFile ID was found.
func (sfq *StagedFileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sfq.Limit(1).IDs(setContextOp(ctx, sfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{stagedfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sfq *StagedFileQuery) FirstIDX(ctx context.Context) int {
	id, err := sfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StagedFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StagedFile entity is found.
// Returns a *NotFoundError when no StagedFile entities are found.
func (sfq *StagedFileQuery) Only(ctx context.Context) (*StagedFile, error) {
	nodes, err := sfq.Limit(2).All(setContextOp(ctx, sfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{stagedfile.Label}
	default:
		return nil, &NotSingularError{stagedfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sfq *StagedFileQuery) OnlyX(ctx context.Context) *StagedFile {
	node, err := sfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StagedFile ID in the query.
// Returns a *NotSingularError when more than one StagedFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (sfq *StagedFileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sfq.Limit(2).IDs(setContextOp(ctx, sfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{stagedfile.Label}
	default:
		err = &NotSingularError{stagedfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sfq *StagedFileQuery) OnlyIDX(ctx context.Context) int {
	id, err := sfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StagedFiles.
func (sfq *StagedFileQuery) All(ctx context.Context) ([]*StagedFile, error) {
	ctx = setContextOp(ctx, sfq.ctx, "All")
	if err := sfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StagedFile, *StagedFileQuery]()
	return withInterceptors[[]*StagedFile](ctx, sfq, qr, sfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sfq *StagedFileQuery) AllX(ctx context.Context) []*StagedFile {
	nodes, err := sfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StagedFile IDs.
func (sfq *StagedFileQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sfq.ctx.Unique == nil && sfq.path != nil {
		sfq.Unique(true)
	}
	ctx = setContextOp(ctx, sfq.ctx, "IDs")
	if err = sfq.Select(stagedfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sfq *StagedFileQuery) IDsX(ctx context.Context) []int {
	ids, err := sfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sfq *StagedFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Count")
	if err := sfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sfq, querierCount[*StagedFileQuery](), sfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sfq *StagedFileQuery) CountX(ctx context.Context) int {
	count, err := sfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sfq *StagedFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Exist")
	switch _, err := sfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sfq *StagedFileQuery) ExistX(ctx context.Context) bool {
	exist, err := sfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StagedFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sfq *StagedFileQuery) Clone() *StagedFileQuery {
	if sfq == nil {
		return nil
	}
	return &StagedFileQuery{
		config:     sfq.config,
		ctx:        sfq.ctx.Clone(),
		order:      append([]stagedfile.OrderOption{}, sfq.order...),
		inters:     append([]Interceptor{}, sfq.inters...),
		predicates: append([]predicate.StagedFile{}, sfq.predicates...),
		// clone intermediate query.
		sql:  sfq.sql.Clone(),
		path: sfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StagedFile.Query().
//		GroupBy(stagedfile.FieldPath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sfq *StagedFileQuery) GroupBy(field string, fields ...string) *StagedFileGroupBy {
	sfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StagedFileGroupBy{build: sfq}
	grbuild.flds = &sfq.ctx.Fields
	grbuild.label = stagedfile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//	}
//
//	client.StagedFile.Query().
//		Select(stagedfile.FieldPath).
//		Scan(ctx, &v)
func (sfq *StagedFileQuery) Select(fields ...string) *StagedFileSelect {
	sfq.ctx.Fields = append(sfq.ctx.Fields, fields...)
	sbuild := &StagedFileSelect{StagedFileQuery: sfq}
	sbuild.label = stagedfile.Label
	sbuild.flds, sbuild.scan = &sfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StagedFileSelect configured with the given aggregations.
func (sfq *StagedFileQuery) Aggregate(fns ...AggregateFunc) *StagedFileSelect {
	return sfq.Select().Aggregate(fns...)
}

func (sfq *StagedFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sfq); err != nil {
				return err
			}
		}
	}
	for _, f := range sfq.ctx.Fields {
		if !stagedfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sfq.path != nil {
		prev, err := sfq.path(ctx)
		if err != nil {
			return err
		}
		sfq.sql = prev
	}
	return nil
}

func (sfq *StagedFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StagedFile, error) {
	var (
		nodes = []*StagedFile{}
		_spec = sfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StagedFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StagedFile{config: sfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sfq *StagedFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sfq.querySpec()
	_spec.Node.Columns = sfq.ctx.Fields
	if len(sfq.ctx.Fields) > 0 {
		_spec.Unique = sfq.ctx.Unique != nil && *sfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sfq.driver, _spec)
}

func (sfq *StagedFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(stagedfile.Table, stagedfile.Columns, sqlgraph.NewFieldSpec(stagedfile.FieldID, field.TypeInt))
	_spec.From = sfq.sql
	if unique := sfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sfq.path != nil {
		_spec.Unique = true
	}
	if fields := sfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stagedfile.FieldID)
		for i := range fields {
			if fields[i] != stagedfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sfq *StagedFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sfq.driver.Dialect())
	t1 := builder.Table(stagedfile.Table)
	columns := sfq.ctx.Fields
	if len(columns) == 0 {
		columns = stagedfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sfq.sql != nil {
		selector = sfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sfq.ctx.Unique != nil && *sfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sfq.predicates {
		p(selector)
	}
	for _, p := range sfq.order {
		p(selector)
	}
	if offset := sfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StagedFileGroupBy is the group-by builder for StagedFile entities.
type StagedFileGroupBy struct {
	selector
	build *StagedFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sfgb *StagedFileGroupBy) Aggregate(fns ...AggregateFunc) *StagedFileGroupBy {
	sfgb.fns = append(sfgb.fns, fns...)
	return sfgb
}

// Scan applies the selector query and scans the result into the given value.
func (sfgb *StagedFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfgb.build.ctx, "GroupBy")
	if err := sfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StagedFileQuery, *StagedFileGroupBy](ctx, sfgb.build, sfgb, sfgb.build.inters, v)
}

func (sfgb *StagedFileGroupBy) sqlScan(ctx context.Context, root *StagedFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sfgb.fns))
	for _, fn := range sfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sfgb.flds)+len(sfgb.fns))
		for _, f := range *sfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StagedFileSelect is the builder for selecting fields of StagedFile entities.
type StagedFileSelect struct {
	*StagedFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sfs *StagedFileSelect) Aggregate(fns ...AggregateFunc) *StagedFileSelect {
	sfs.fns = append(sfs.fns, fns...)
	return sfs
}

// Scan applies the selector query and scans the result into the given value.
func (sfs *StagedFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfs.ctx, "Select")
	if err := sfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StagedFileQuery, *StagedFileSelect](ctx, sfs.StagedFileQuery, sfs, sfs.inters, v)
}

func (sfs *StagedFileSelect) sqlScan(ctx context.Context, root *StagedFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sfs.fns))
	for _, fn := range sfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
