// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Creaft-JP/tit/db/global/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/global/ent/globalconfig"
	"github.com/Creaft-JP/tit/db/global/ent/logintoken"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// GlobalConfig is the client for interacting with the GlobalConfig builders.
	GlobalConfig *GlobalConfigClient
	// LoginToken is the client for interacting with the LoginToken builders.
	LoginToken *LoginTokenClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.GlobalConfig = NewGlobalConfigClient(c.config)
	c.LoginToken = NewLoginTokenClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		GlobalConfig: NewGlobalConfigClient(cfg),
		LoginToken:   NewLoginTokenClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		GlobalConfig: NewGlobalConfigClient(cfg),
		LoginToken:   NewLoginTokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		GlobalConfig.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.GlobalConfig.Use(hooks...)
	c.LoginToken.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.GlobalConfig.Intercept(interceptors...)
	c.LoginToken.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *GlobalConfigMutation:
		return c.GlobalConfig.mutate(ctx, m)
	case *LoginTokenMutation:
		return c.LoginToken.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// GlobalConfigClient is a client for the GlobalConfig schema.
type GlobalConfigClient struct {
	config
}

// NewGlobalConfigClient returns a client for the GlobalConfig from the given config.
func NewGlobalConfigClient(c config) *GlobalConfigClient {
	return &GlobalConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `globalconfig.Hooks(f(g(h())))`.
func (c *GlobalConfigClient) Use(hooks ...Hook) {
	c.hooks.GlobalConfig = append(c.hooks.GlobalConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `globalconfig.Intercept(f(g(h())))`.
func (c *GlobalConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.GlobalConfig = append(c.inters.GlobalConfig, interceptors...)
}

// Create returns a builder for creating a GlobalConfig entity.
func (c *GlobalConfigClient) Create() *GlobalConfigCreate {
	mutation := newGlobalConfigMutation(c.config, OpCreate)
	return &GlobalConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GlobalConfig entities.
func (c *GlobalConfigClient) CreateBulk(builders ...*GlobalConfigCreate) *GlobalConfigCreateBulk {
	return &GlobalConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GlobalConfig.
func (c *GlobalConfigClient) Update() *GlobalConfigUpdate {
	mutation := newGlobalConfigMutation(c.config, OpUpdate)
	return &GlobalConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GlobalConfigClient) UpdateOne(gc *GlobalConfig) *GlobalConfigUpdateOne {
	mutation := newGlobalConfigMutation(c.config, OpUpdateOne, withGlobalConfig(gc))
	return &GlobalConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GlobalConfigClient) UpdateOneID(id int) *GlobalConfigUpdateOne {
	mutation := newGlobalConfigMutation(c.config, OpUpdateOne, withGlobalConfigID(id))
	return &GlobalConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GlobalConfig.
func (c *GlobalConfigClient) Delete() *GlobalConfigDelete {
	mutation := newGlobalConfigMutation(c.config, OpDelete)
	return &GlobalConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GlobalConfigClient) DeleteOne(gc *GlobalConfig) *GlobalConfigDeleteOne {
	return c.DeleteOneID(gc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GlobalConfigClient) DeleteOneID(id int) *GlobalConfigDeleteOne {
	builder := c.Delete().Where(globalconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GlobalConfigDeleteOne{builder}
}

// Query returns a query builder for GlobalConfig.
func (c *GlobalConfigClient) Query() *GlobalConfigQuery {
	return &GlobalConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGlobalConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a GlobalConfig entity by its id.
func (c *GlobalConfigClient) Get(ctx context.Context, id int) (*GlobalConfig, error) {
	return c.Query().Where(globalconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GlobalConfigClient) GetX(ctx context.Context, id int) *GlobalConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GlobalConfigClient) Hooks() []Hook {
	return c.hooks.GlobalConfig
}

// Interceptors returns the client interceptors.
func (c *GlobalConfigClient) Interceptors() []Interceptor {
	return c.inters.GlobalConfig
}

func (c *GlobalConfigClient) mutate(ctx context.Context, m *GlobalConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GlobalConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GlobalConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GlobalConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GlobalConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown GlobalConfig mutation op: %q", m.Op())
	}
}

// LoginTokenClient is a client for the LoginToken schema.
type LoginTokenClient struct {
	config
}

// NewLoginTokenClient returns a client for the LoginToken from the given config.
func NewLoginTokenClient(c config) *LoginTokenClient {
	return &LoginTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `logintoken.Hooks(f(g(h())))`.
func (c *LoginTokenClient) Use(hooks ...Hook) {
	c.hooks.LoginToken = append(c.hooks.LoginToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `logintoken.Intercept(f(g(h())))`.
func (c *LoginTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.LoginToken = append(c.inters.LoginToken, interceptors...)
}

// Create returns a builder for creating a LoginToken entity.
func (c *LoginTokenClient) Create() *LoginTokenCreate {
	mutation := newLoginTokenMutation(c.config, OpCreate)
	return &LoginTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LoginToken entities.
func (c *LoginTokenClient) CreateBulk(builders ...*LoginTokenCreate) *LoginTokenCreateBulk {
	return &LoginTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LoginToken.
func (c *LoginTokenClient) Update() *LoginTokenUpdate {
	mutation := newLoginTokenMutation(c.config, OpUpdate)
	return &LoginTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LoginTokenClient) UpdateOne(lt *LoginToken) *LoginTokenUpdateOne {
	mutation := newLoginTokenMutation(c.config, OpUpdateOne, withLoginToken(lt))
	return &LoginTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LoginTokenClient) UpdateOneID(id int) *LoginTokenUpdateOne {
	mutation := newLoginTokenMutation(c.config, OpUpdateOne, withLoginTokenID(id))
	return &LoginTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LoginToken.
func (c *LoginTokenClient) Delete() *LoginTokenDelete {
	mutation := newLoginTokenMutation(c.config, OpDelete)
	return &LoginTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LoginTokenClient) DeleteOne(lt *LoginToken) *LoginTokenDeleteOne {
	return c.DeleteOneID(lt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LoginTokenClient) DeleteOneID(id int) *LoginTokenDeleteOne {
	builder := c.Delete().Where(logintoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LoginTokenDeleteOne{builder}
}

// Query returns a query builder for LoginToken.
func (c *LoginTokenClient) Query() *LoginTokenQuery {
	return &LoginTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLoginToken},
		inters: c.Interceptors(),
	}
}

// Get returns a LoginToken entity by its id.
func (c *LoginTokenClient) Get(ctx context.Context, id int) (*LoginToken, error) {
	return c.Query().Where(logintoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LoginTokenClient) GetX(ctx context.Context, id int) *LoginToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LoginTokenClient) Hooks() []Hook {
	return c.hooks.LoginToken
}

// Interceptors returns the client interceptors.
func (c *LoginTokenClient) Interceptors() []Interceptor {
	return c.inters.LoginToken
}

func (c *LoginTokenClient) mutate(ctx context.Context, m *LoginTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LoginTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LoginTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LoginTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LoginTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LoginToken mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		GlobalConfig, LoginToken []ent.Hook
	}
	inters struct {
		GlobalConfig, LoginToken []ent.Interceptor
	}
)