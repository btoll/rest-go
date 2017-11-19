package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamOpeningConfigModel struct{}

func (m *TeamOpeningConfigModel) AllocateID(ctx *Context) error {
	return AllocateSequentialID(ctx)
}

func (m *TeamOpeningConfigModel) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateTeamOpeningConfigContext:
		return &Context{
			Kind:    constants.DB_TEAM_OPENING_CONFIG,
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteTeamOpeningConfigContext:
		return &Context{
			Kind:   constants.DB_TEAM_OPENING_CONFIG,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListTeamOpeningConfigContext:
		return &Context{
			Kind:   constants.DB_TEAM_OPENING_CONFIG,
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowTeamOpeningConfigContext:
		return &Context{
			Kind:   constants.DB_TEAM_OPENING_CONFIG,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateTeamOpeningConfigContext:
		return &Context{
			Kind:    constants.DB_TEAM_OPENING_CONFIG,
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *TeamOpeningConfigModel) GetModelInstance() interface{} {
	return &app.TeamOpeningConfigMedia{}
}

func (m *TeamOpeningConfigModel) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := app.TeamOpeningConfigMediaCollection{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *TeamOpeningConfigModel) SetModel(ctx *Context) error {
	rec := &app.TeamOpeningConfigMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, GetKey(ctx), rec)

	return err
}
