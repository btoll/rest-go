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
		// TODO: Return error.
		return &Context{}
	}
}

func (m *TeamPersist) GetAll(q *datastore.Query, mc *Context) *QueryResponse {
	c := []TeamPersist{}
	_, err := q.GetAll(mc.GaeCtx, &c)

	ms := make(app.TeamMediaCollection, len(c))

	for i, model := range c {
		tm := &app.TeamMedia{}
		copier.Copy(tm, model)
		ms[i] = tm
	}

	return &QueryResponse{
		Models: &ms,
		Err:    err,
	}
}

func (m *TeamPersist) GetOne(q *datastore.Query, mc *Context) *QueryResponse {
	var tp *TeamPersist
	var k *datastore.Key
	var err error

	q = q.KeysOnly()
	for iter := q.Run(mc.GaeCtx); ; {
		var x TeamPersist

		key, err := iter.Next(&x)

		if err == datastore.Done {
			break
		}

		tp = &TeamPersist{}
		err = datastore.Get(mc.GaeCtx, key, tp)
		k = key
	}

	return &QueryResponse{
		Key:    k,
		Models: tp,
		Err:    err,
	}
}

func (m *TeamPersist) Add(ctx *Context) error {
	return m.Set(ctx, datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil))
}

func (m *TeamPersist) Set(ctx *Context, key *datastore.Key) error {
	tp := &app.TeamPayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return err
}
