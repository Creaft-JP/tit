// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/committedfile"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
	"github.com/Creaft-JP/tit/db/local/ent/section"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
)

// TitCommitQuery is the builder for querying TitCommit entities.
type TitCommitQuery struct {
	config
	ctx         *QueryContext
	order       []titcommit.OrderOption
	inters      []Interceptor
	predicates  []predicate.TitCommit
	withSection *SectionQuery
	withFiles   *CommittedFileQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TitCommitQuery builder.
func (tcq *TitCommitQuery) Where(ps ...predicate.TitCommit) *TitCommitQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit the number of records to be returned by this query.
func (tcq *TitCommitQuery) Limit(limit int) *TitCommitQuery {
	tcq.ctx.Limit = &limit
	return tcq
}

// Offset to start from.
func (tcq *TitCommitQuery) Offset(offset int) *TitCommitQuery {
	tcq.ctx.Offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TitCommitQuery) Unique(unique bool) *TitCommitQuery {
	tcq.ctx.Unique = &unique
	return tcq
}

// Order specifies how the records should be ordered.
func (tcq *TitCommitQuery) Order(o ...titcommit.OrderOption) *TitCommitQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QuerySection chains the current query on the "section" edge.
func (tcq *TitCommitQuery) QuerySection() *SectionQuery {
	query := (&SectionClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(titcommit.Table, titcommit.FieldID, selector),
			sqlgraph.To(section.Table, section.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, titcommit.SectionTable, titcommit.SectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (tcq *TitCommitQuery) QueryFiles() *CommittedFileQuery {
	query := (&CommittedFileClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(titcommit.Table, titcommit.FieldID, selector),
			sqlgraph.To(committedfile.Table, committedfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, titcommit.FilesTable, titcommit.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TitCommit entity from the query.
// Returns a *NotFoundError when no TitCommit was found.
func (tcq *TitCommitQuery) First(ctx context.Context) (*TitCommit, error) {
	nodes, err := tcq.Limit(1).All(setContextOp(ctx, tcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{titcommit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TitCommitQuery) FirstX(ctx context.Context) *TitCommit {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TitCommit ID from the query.
// Returns a *NotFoundError when no TitCommit ID was found.
func (tcq *TitCommitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(setContextOp(ctx, tcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{titcommit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TitCommitQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TitCommit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TitCommit entity is found.
// Returns a *NotFoundError when no TitCommit entities are found.
func (tcq *TitCommitQuery) Only(ctx context.Context) (*TitCommit, error) {
	nodes, err := tcq.Limit(2).All(setContextOp(ctx, tcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{titcommit.Label}
	default:
		return nil, &NotSingularError{titcommit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TitCommitQuery) OnlyX(ctx context.Context) *TitCommit {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TitCommit ID in the query.
// Returns a *NotSingularError when more than one TitCommit ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TitCommitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(setContextOp(ctx, tcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{titcommit.Label}
	default:
		err = &NotSingularError{titcommit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TitCommitQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TitCommits.
func (tcq *TitCommitQuery) All(ctx context.Context) ([]*TitCommit, error) {
	ctx = setContextOp(ctx, tcq.ctx, "All")
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TitCommit, *TitCommitQuery]()
	return withInterceptors[[]*TitCommit](ctx, tcq, qr, tcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TitCommitQuery) AllX(ctx context.Context) []*TitCommit {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TitCommit IDs.
func (tcq *TitCommitQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tcq.ctx.Unique == nil && tcq.path != nil {
		tcq.Unique(true)
	}
	ctx = setContextOp(ctx, tcq.ctx, "IDs")
	if err = tcq.Select(titcommit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TitCommitQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TitCommitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Count")
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcq, querierCount[*TitCommitQuery](), tcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TitCommitQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TitCommitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Exist")
	switch _, err := tcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TitCommitQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TitCommitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TitCommitQuery) Clone() *TitCommitQuery {
	if tcq == nil {
		return nil
	}
	return &TitCommitQuery{
		config:      tcq.config,
		ctx:         tcq.ctx.Clone(),
		order:       append([]titcommit.OrderOption{}, tcq.order...),
		inters:      append([]Interceptor{}, tcq.inters...),
		predicates:  append([]predicate.TitCommit{}, tcq.predicates...),
		withSection: tcq.withSection.Clone(),
		withFiles:   tcq.withFiles.Clone(),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
	}
}

// WithSection tells the query-builder to eager-load the nodes that are connected to
// the "section" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TitCommitQuery) WithSection(opts ...func(*SectionQuery)) *TitCommitQuery {
	query := (&SectionClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withSection = query
	return tcq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TitCommitQuery) WithFiles(opts ...func(*CommittedFileQuery)) *TitCommitQuery {
	query := (&CommittedFileClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withFiles = query
	return tcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Number int `json:"number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TitCommit.Query().
//		GroupBy(titcommit.FieldNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *TitCommitQuery) GroupBy(field string, fields ...string) *TitCommitGroupBy {
	tcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TitCommitGroupBy{build: tcq}
	grbuild.flds = &tcq.ctx.Fields
	grbuild.label = titcommit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Number int `json:"number,omitempty"`
//	}
//
//	client.TitCommit.Query().
//		Select(titcommit.FieldNumber).
//		Scan(ctx, &v)
func (tcq *TitCommitQuery) Select(fields ...string) *TitCommitSelect {
	tcq.ctx.Fields = append(tcq.ctx.Fields, fields...)
	sbuild := &TitCommitSelect{TitCommitQuery: tcq}
	sbuild.label = titcommit.Label
	sbuild.flds, sbuild.scan = &tcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TitCommitSelect configured with the given aggregations.
func (tcq *TitCommitQuery) Aggregate(fns ...AggregateFunc) *TitCommitSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *TitCommitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tcq); err != nil {
				return err
			}
		}
	}
	for _, f := range tcq.ctx.Fields {
		if !titcommit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TitCommitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TitCommit, error) {
	var (
		nodes       = []*TitCommit{}
		withFKs     = tcq.withFKs
		_spec       = tcq.querySpec()
		loadedTypes = [2]bool{
			tcq.withSection != nil,
			tcq.withFiles != nil,
		}
	)
	if tcq.withSection != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, titcommit.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TitCommit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TitCommit{config: tcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tcq.withSection; query != nil {
		if err := tcq.loadSection(ctx, query, nodes, nil,
			func(n *TitCommit, e *Section) { n.Edges.Section = e }); err != nil {
			return nil, err
		}
	}
	if query := tcq.withFiles; query != nil {
		if err := tcq.loadFiles(ctx, query, nodes,
			func(n *TitCommit) { n.Edges.Files = []*CommittedFile{} },
			func(n *TitCommit, e *CommittedFile) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tcq *TitCommitQuery) loadSection(ctx context.Context, query *SectionQuery, nodes []*TitCommit, init func(*TitCommit), assign func(*TitCommit, *Section)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TitCommit)
	for i := range nodes {
		if nodes[i].section_commits == nil {
			continue
		}
		fk := *nodes[i].section_commits
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(section.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "section_commits" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tcq *TitCommitQuery) loadFiles(ctx context.Context, query *CommittedFileQuery, nodes []*TitCommit, init func(*TitCommit), assign func(*TitCommit, *CommittedFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TitCommit)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CommittedFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(titcommit.FilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.tit_commit_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "tit_commit_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tit_commit_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tcq *TitCommitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.ctx.Fields
	if len(tcq.ctx.Fields) > 0 {
		_spec.Unique = tcq.ctx.Unique != nil && *tcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TitCommitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(titcommit.Table, titcommit.Columns, sqlgraph.NewFieldSpec(titcommit.FieldID, field.TypeInt))
	_spec.From = tcq.sql
	if unique := tcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcq.path != nil {
		_spec.Unique = true
	}
	if fields := tcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, titcommit.FieldID)
		for i := range fields {
			if fields[i] != titcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TitCommitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(titcommit.Table)
	columns := tcq.ctx.Fields
	if len(columns) == 0 {
		columns = titcommit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.ctx.Unique != nil && *tcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TitCommitGroupBy is the group-by builder for TitCommit entities.
type TitCommitGroupBy struct {
	selector
	build *TitCommitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TitCommitGroupBy) Aggregate(fns ...AggregateFunc) *TitCommitGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tcgb *TitCommitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcgb.build.ctx, "GroupBy")
	if err := tcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TitCommitQuery, *TitCommitGroupBy](ctx, tcgb.build, tcgb, tcgb.build.inters, v)
}

func (tcgb *TitCommitGroupBy) sqlScan(ctx context.Context, root *TitCommitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tcgb.flds)+len(tcgb.fns))
		for _, f := range *tcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TitCommitSelect is the builder for selecting fields of TitCommit entities.
type TitCommitSelect struct {
	*TitCommitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *TitCommitSelect) Aggregate(fns ...AggregateFunc) *TitCommitSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TitCommitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcs.ctx, "Select")
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TitCommitQuery, *TitCommitSelect](ctx, tcs.TitCommitQuery, tcs, tcs.inters, v)
}

func (tcs *TitCommitSelect) sqlScan(ctx context.Context, root *TitCommitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tcs.fns))
	for _, fn := range tcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
