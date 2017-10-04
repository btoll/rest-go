package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type SportPersist struct {
	app.SportPayload
	Context
}

func (m *SportPersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateSportContext:
		return &Context{
			Kind:    "SportPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteSportContext:
		return &Context{
			Kind:   "SportPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListSportContext:
		return &Context{
			Kind:   "SportPersist",
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowSportContext:
		return &Context{
			Kind:   "SportPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateSportContext:
		return &Context{
			Kind:    "SportPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *SportPersist) GetModel() interface{} {
	return &app.SportMedia{}
}

func (m *SportPersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.SportMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *SportPersist) SetModel(ctx *Context, key *datastore.Key) error {
	rec := &app.SportMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return err
}
