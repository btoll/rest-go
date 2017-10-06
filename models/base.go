package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
	"google.golang.org/appengine/datastore"
)

type App interface {
	Create() (string, error)
	Read() (Model, error)
	Update() error
	Delete() error
	List() ([]string, []Model, error)
}

type Model interface {
	GetCtx(ctx context.Context) *Context
	GetModelCollection(c *Context) ([]*datastore.Key, interface{}, error)
	GetModel() interface{}
	SetModel(ctx *Context, key *datastore.Key) error
}

type Context struct {
	GaeCtx  context.Context
	Payload interface{}
	Kind    string
	ID      string
	App
}

type ModelStore struct {
	Keys   []string    `json:"keys"`
	Models interface{} `json:"models"`
}

const (
	GAME                = constants.DB_GAME
	EVENT               = constants.DB_EVENT
	SPORT               = constants.DB_SPORT
	TEAM                = constants.DB_TEAM
	TEAM_OPENING_CONFIG = constants.DB_TEAM_OPENING_CONFIG
)

func ModelFactory(name string) Model {
	switch name {
	case GAME:
		return new(GamePersist)
	case EVENT:
		return new(EventPersist)
	case SPORT:
		return new(SportPersist)
	case TEAM:
		return new(TeamPersist)
	case TEAM_OPENING_CONFIG:
		return new(TeamOpeningConfigPersist)
	}

	// TODO: Return error?
	return nil
}

func GetCtx(kind string, ctx context.Context) *Context {
	return ModelFactory(kind).GetCtx(ctx)
}

func (ctx *Context) getKey() *datastore.Key {
	return datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil)
}

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func (ctx *Context) Create() (string, error) {
	// Returns an int64. We only need one ID returned.
	low, _, err := datastore.AllocateIDs(ctx.GaeCtx, ctx.Kind, nil, 1)

	if err != nil {
		return "", goa.ErrInternal(err, "endpoint", "create")
	}

	ctx.ID = strconv.FormatInt(low, 10)
	err = ModelFactory(ctx.Kind).SetModel(ctx, ctx.getKey())

	if err != nil {
		return "", goa.ErrInternal(err, "endpoint", "create")
	}

	return ctx.ID, nil
}

func (ctx *Context) Read() (interface{}, error) {
	model := ModelFactory(ctx.Kind).GetModel()

	if err := datastore.Get(ctx.GaeCtx, ctx.getKey(), model); err != nil {
		return nil, goa.ErrBadRequest(err, "endpoint", "show")
	}

	return model, nil
}

func (ctx *Context) Update() error {
	err := ModelFactory(ctx.Kind).SetModel(ctx, ctx.getKey())

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "update")
	}

	return nil
}

func (ctx *Context) Delete() error {
	err := datastore.Delete(ctx.GaeCtx, ctx.getKey())

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "delete")
	}

	return nil
}

func (ctx *Context) List() ([]byte, error) {
	keys, c, err := ModelFactory(ctx.Kind).GetModelCollection(ctx)

	if err != nil {
		return nil, goa.ErrInternal(err, "endpoint", "list")
	}

	store := &ModelStore{
		Keys:   make([]string, len(keys)),
		Models: c,
	}

	for i, key := range keys {
		store.Keys[i] = key.StringID()
	}

	b, err := json.Marshal(store)

	if err != nil {
		return nil, goa.ErrInternal(err, "endpoint", "list")
	}

	return b, nil
}
