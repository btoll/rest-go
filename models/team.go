package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamPersist struct{}

func (m *TeamPersist) AllocateID(ctx *Context) error {
	return AllocateSequentialID(ctx)
}

func (m *TeamPersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateTeamContext:
		return &Context{
			Kind:    constants.DB_TEAM,
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteTeamContext:
		return &Context{
			Kind:   constants.DB_TEAM,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListTeamContext:
		return &Context{
			Kind:   constants.DB_TEAM,
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowTeamContext:
		return &Context{
			Kind:   constants.DB_TEAM,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateTeamContext:
		return &Context{
			Kind:    constants.DB_TEAM,
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *TeamPersist) GetModelInstance() interface{} {
	return &app.TeamMedia{}
}

func (m *TeamPersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.TeamMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *TeamPersist) SetModel(ctx *Context) error {
	rec := &app.TeamMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, GetKey(ctx), rec)

	return err
}
