package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

type App interface {
	Create() (string, error)
	Read() (Model, error)
	Update() error
	Delete() error
	List() ([]Model, error)
}

type Model interface {
	GetCtx(ctx context.Context) *Context
	GetModel() interface{}
	GetModelCollection(c *Context) (interface{}, error)
	Set(ctx *Context, key *datastore.Key) (string, error)
}

type Context struct {
	GaeCtx  context.Context
	Payload interface{}
	Kind    string
	ID      string // serialized datastore.Key
	App
}

const (
	GAME                = "GamePersist"
	EVENT               = "EventPersist"
	SPORT               = "SportPersist"
	TEAM                = "TeamPersist"
	TEAM_OPENING_CONFIG = "TeamOpeningConfigPersist"
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

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func (ctx *Context) Create() (string, error) {
	key := datastore.NewIncompleteKey(ctx.GaeCtx, ctx.Kind, nil)
	id, err := ModelFactory(ctx.Kind).Set(ctx, key)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (ctx *Context) Read() (interface{}, error) {
	decodedKey, err := datastore.DecodeKey(ctx.ID)

	if err != nil {
		return nil, err
	}

	model := ModelFactory(ctx.Kind).GetModel()

	if datastore.Get(ctx.GaeCtx, decodedKey, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (ctx *Context) Update() error {
	decodedKey, err := datastore.DecodeKey(ctx.ID)

	if err != nil {
		return err
	}

	_, err = ModelFactory(ctx.Kind).Set(ctx, decodedKey)

	return err
}

func (ctx *Context) Delete() error {
	decodedKey, err := datastore.DecodeKey(ctx.ID)

	if err != nil {
		return err
	}

	return datastore.Delete(ctx.GaeCtx, decodedKey)
}

func (ctx *Context) List() (interface{}, error) {
	c, err := ModelFactory(ctx.Kind).GetModelCollection(ctx)

	if err != nil {
		return nil, err
	}

	return c, nil
}

/* -----------------------------------------------------------
    ERROR (TODO)
----------------------------------------------------------- */
