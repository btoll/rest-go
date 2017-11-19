package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type GameModel struct{}

func (m *GameModel) AllocateID(ctx *Context) error {
	return AllocateSequentialID(ctx)
}

func (m *GameModel) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateGameContext:
		return &Context{
			Kind:    constants.DB_GAME,
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateGameContext:
		return &Context{
			Kind:    constants.DB_GAME,
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *GameModel) GetModelInstance() interface{} {
	return &app.GameMedia{}
}

func (m *GameModel) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := app.GameMediaCollection{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *GameModel) SetModel(ctx *Context) error {
	rec := &app.GameMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, GetKey(ctx), rec)

	return err
}
