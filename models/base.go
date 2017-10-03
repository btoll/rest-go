package models

import (
	"context"
	"strconv"

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
	Keys   []int64     `json:"keys"`
	Models interface{} `json:"models"`
}

const (
	GAME                = "GamePersist"
	EVENT               = "EventPersist"
	SPORT               = "SportPersist"
	TEAM                = "TeamPersist"
	TEAM_OPENING_CONFIG = "TeamOpeningConfigPersist"
)

func getKey(ctx *Context) *datastore.Key {
	// This should never error b/c we're guaranteeing that ctx.ID is a
	// valid string :)
	id, _ := strconv.ParseInt(ctx.ID, 10, 64)
	return datastore.NewKey(ctx.GaeCtx, ctx.Kind, "", id, nil)
}

func ModelFactory(name string) Model {
	switch name {
	//	case GAME:
	//		return new(GamePersist)
	//	case EVENT:
	//		return new(EventPersist)
	//	case SPORT:
	//		return new(SportPersist)
	case TEAM:
		return new(TeamPersist)
		//	case TEAM_OPENING_CONFIG:
		//		return new(TeamOpeningConfigPersist)
	}

	// TODO: Return error?
	return nil
}

func GetCtx(kind string, ctx context.Context) *Context {
	return ModelFactory(kind).GetCtx(ctx)
}

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func (ctx *Context) Create() (string, error) {
	// Returns an int64.
	low, _, err := datastore.AllocateIDs(ctx.GaeCtx, ctx.Kind, nil, 1)

	if err != nil {
		return "", err
	}

	key := datastore.NewKey(ctx.GaeCtx, ctx.Kind, "", low, nil)
	err = ModelFactory(ctx.Kind).SetModel(ctx, key)

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(low, 10), nil
}

func (ctx *Context) Read() (interface{}, error) {
	model := ModelFactory(ctx.Kind).GetModel()

	if err := datastore.Get(ctx.GaeCtx, getKey(ctx), model); err != nil {
		return nil, err
	}

	return model, nil
}

func (ctx *Context) Update() error {
	return ModelFactory(ctx.Kind).SetModel(ctx, getKey(ctx))
}

func (ctx *Context) Delete() error {
	return datastore.Delete(ctx.GaeCtx, getKey(ctx))
}

func (ctx *Context) List() (*ModelStore, error) {
	keys, c, err := ModelFactory(ctx.Kind).GetModelCollection(ctx)

	if err != nil {
		return nil, err
	}

	store := &ModelStore{
		Keys:   make([]int64, len(keys)),
		Models: c,
	}

	for i, key := range keys {
		store.Keys[i] = key.IntID()
	}

	if err != nil {
		return nil, err
	}

	return store, nil
}

/* -----------------------------------------------------------
    ERROR (TODO)
----------------------------------------------------------- */
