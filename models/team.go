package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamPersist struct {
	app.TeamPayload
	Context
}

func (m *TeamPersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateTeamContext:
		return &Context{
			Kind:    "TeamPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteTeamContext:
		return &Context{
			Kind:   "TeamPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListTeamContext:
		return &Context{
			Kind:   "TeamPersist",
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowTeamContext:
		return &Context{
			Kind:   "TeamPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateTeamContext:
		return &Context{
			Kind:    "TeamPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *TeamPersist) GetModel() interface{} {
	return &app.TeamMedia{}
}

func (m *TeamPersist) GetModelCollection(ctx *Context) (interface{}, error) {
	c := app.TeamMediaCollection{}

	if _, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c); err != nil {
		return nil, err
	}

	return c, nil
}

func (m *TeamPersist) Set(ctx *Context, key *datastore.Key) (string, error) {
	rec := &app.TeamMedia{}

	copier.Copy(rec, ctx.Payload)
	newkey, err := datastore.Put(ctx.GaeCtx, key, rec)

	if err != nil {
		return "", err
	}

	return newkey.Encode(), nil
}
