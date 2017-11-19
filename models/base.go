package models

import (
	"context"
	"strconv"

	"github.com/dgaedcke/nmg_shared/constants"
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

func initModel(kind string, c context.Context) (Model, *Context) {
	model := modelFactory(kind)
	return model, model.GetCtx(c)
}

func modelFactory(name string) Model {
	switch name {
	case GAME:
		return new(GameModel)
	case EVENT:
		return new(EventModel)
	case SPORT:
		return new(SportModel)
	case TEAM:
		return new(TeamModel)
	case TEAM_OPENING_CONFIG:
		return new(TeamOpeningConfigModel)
	}

	// TODO: Return error?
	return nil
}

// Will be used by all models that don't need a custom id like "baseball".
func AllocateSequentialID(ctx *Context) error {
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

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func Create(kind string, c context.Context) (string, error) {
	model, ctx := initModel(kind, c)
	// Get unique string ID, will be different by model.
	err := model.AllocateID(ctx)
	if err != nil {
		return "", err
	}
	err = model.SetModel(ctx)
	if err != nil {
		return "", err
	}
	return ctx.ID, nil
}

func Read(kind string, c context.Context) (interface{}, error) {
	model, ctx := initModel(kind, c)
	instance := model.GetModelInstance()
	if err := datastore.Get(ctx.GaeCtx, GetKey(ctx), instance); err != nil {
		return nil, err
	}
	return instance, nil
}

func Update(kind string, c context.Context) error {
	model, ctx := initModel(kind, c)
	err := model.SetModel(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Delete(kind string, c context.Context) error {
	_, ctx := initModel(kind, c)
	err := datastore.Delete(ctx.GaeCtx, GetKey(ctx))
	if err != nil {
		return err
	}
	return nil
}

func List(kind string, c context.Context) (interface{}, error) {
	model, ctx := initModel(kind, c)
	_, coll, err := model.GetModelCollection(ctx)
	if err != nil {
		return nil, err
	}
	return coll, nil
}
