package models

import (
	"context"

	"github.com/dgaedcke/nmg_admin_service/app"
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

func (m *TeamOpeningConfigPersist) GetAll(q *datastore.Query, mc *Context) *QueryResponse {
	c := []TeamOpeningConfigPersist{}
	_, err := q.GetAll(mc.GaeCtx, &c)

	ms := make(app.TeamOpeningConfigMediaCollection, len(c))

	for i, model := range c {
		tm := &app.TeamOpeningConfigMedia{}
		copier.Copy(tm, model)
		ms[i] = tm
	}

	return &QueryResponse{
		Models: &ms,
		Err:    err,
	}
}

func (m *TeamOpeningConfigPersist) GetOne(q *datastore.Query, mc *Context) *QueryResponse {
	var tp *TeamOpeningConfigPersist
	var k *datastore.Key
	var err error

	q = q.KeysOnly()
	for iter := q.Run(mc.GaeCtx); ; {
		var x TeamOpeningConfigPersist

		key, err := iter.Next(&x)

		if err == datastore.Done {
			break
		}

		tp = &TeamOpeningConfigPersist{}
		err = datastore.Get(mc.GaeCtx, key, tp)
		k = key
	}

	return &QueryResponse{
		Key:    k,
		Models: tp,
		Err:    err,
	}
}

func (m *TeamOpeningConfigPersist) Add(ctx *Context) error {
	return m.Set(ctx, datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil))
}

func (m *TeamOpeningConfigPersist) Set(ctx *Context, key *datastore.Key) error {
	tp := &app.TeamOpeningConfigPayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return err
}
