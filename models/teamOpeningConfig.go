package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamOpeningConfigPersist struct {
	app.TeamOpeningConfigPayload
	Context
}

func (m *TeamOpeningConfigPersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateTeamOpeningConfigContext:
		return &Context{
			Kind:    "TeamOpeningConfigPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteTeamOpeningConfigContext:
		return &Context{
			Kind:   "TeamOpeningConfigPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListTeamOpeningConfigContext:
		return &Context{
			Kind:   "TeamOpeningConfigPersist",
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowTeamOpeningConfigContext:
		return &Context{
			Kind:   "TeamOpeningConfigPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateTeamOpeningConfigContext:
		return &Context{
			Kind:    "TeamOpeningConfigPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error.
		return &Context{}
	}
}

func (m *TeamOpeningConfigPersist) GetModelCollection(ctx *Context) (interface{}, error) {
	return nil, nil
}

func (m *TeamOpeningConfigPersist) GetModel() interface{} {
	return &app.TeamOpeningConfigMedia{}
}

func (m *TeamOpeningConfigPersist) Set(ctx *Context, key *datastore.Key) (string, error) {
	tp := &app.TeamOpeningConfigPayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return "", err
}
