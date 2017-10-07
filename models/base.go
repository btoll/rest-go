package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
	"google.golang.org/appengine/datastore"
)

const (
	GAME                = constants.DB_GAME
	EVENT               = constants.DB_EVENT
	SPORT               = constants.DB_SPORT
	TEAM                = constants.DB_TEAM
	TEAM_OPENING_CONFIG = constants.DB_TEAM_OPENING_CONFIG
)

type Model interface {
	AllocateID(ctx *Context) error
	GetCtx(ctx context.Context) *Context
	GetModelCollection(c *Context) ([]*datastore.Key, interface{}, error)
	GetModelInstance() interface{}
	SetModel(ctx *Context) error
}

type Context struct {
	GaeCtx  context.Context
	Payload interface{}
	Kind    string
	ID      string
}

type ModelStore struct {
	Keys   []string    `json:"keys"`
	Models interface{} `json:"models"`
}

func initModel(kind string, c context.Context) (Model, *Context) {
	model := ModelFactory(kind)
	return model, model.GetCtx(c)
}

// This function will be used by all models that don't need a custom id
// like "baseball".
func Allocate(ctx *Context) error {
	// Returns an int64. We only need one ID returned.
	low, _, err := datastore.AllocateIDs(ctx.GaeCtx, ctx.Kind, nil, 1)

	if err != nil {
		return err
	}

	ctx.ID = strconv.FormatInt(low, 10)

	return nil
}

func GetKey(ctx *Context) *datastore.Key {
	return datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil)
}

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

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func Create(kind string, c context.Context) (string, error) {
	model, ctx := initModel(kind, c)

	// Get unique string ID, will be different by model.
	err := model.AllocateID(ctx)
	if err != nil {
		return "", goa.ErrInternal(err, "endpoint", "create")
	}

	err = model.SetModel(ctx)
	if err != nil {
		return "", goa.ErrInternal(err, "endpoint", "create")
	}

	return ctx.ID, nil
}

func Read(kind string, c context.Context) (interface{}, error) {
	model, ctx := initModel(kind, c)
	instance := model.GetModelInstance()

	if err := datastore.Get(ctx.GaeCtx, GetKey(ctx), instance); err != nil {
		return nil, goa.ErrBadRequest(err, "endpoint", "show")
	}

	return instance, nil
}

func Update(kind string, c context.Context) error {
	model, ctx := initModel(kind, c)

	err := model.SetModel(ctx)
	if err != nil {
		return goa.ErrInternal(err, "endpoint", "update")
	}

	return nil
}

func Delete(kind string, c context.Context) error {
	_, ctx := initModel(kind, c)

	err := datastore.Delete(ctx.GaeCtx, GetKey(ctx))
	if err != nil {
		return goa.ErrInternal(err, "endpoint", "delete")
	}

	return nil
}

func List(kind string, c context.Context) ([]byte, error) {
	model, ctx := initModel(kind, c)

	keys, coll, err := model.GetModelCollection(ctx)
	if err != nil {
		return nil, goa.ErrInternal(err, "endpoint", "list")
	}

	store := &ModelStore{
		Keys:   make([]string, len(keys)),
		Models: coll,
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
