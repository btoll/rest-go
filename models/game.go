package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type GamePersist struct {
	app.GamePayload
	Context
}

func (m *GamePersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateGameContext:
		return &Context{
			Kind:    "GamePersist",
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteGameContext:
		return &Context{
			Kind:   "GamePersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListGameContext:
		return &Context{
			Kind:   "GamePersist",
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowGameContext:
		return &Context{
			Kind:   "GamePersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateGameContext:
		return &Context{
			Kind:    "GamePersist",
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *GamePersist) GetModel() interface{} {
	return &app.GameMedia{}
}

func (m *GamePersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.GameMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *GamePersist) SetModel(ctx *Context, key *datastore.Key) error {
	rec := &app.GameMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return err
}
