package models

import (
	"context"

	"github.com/dgaedcke/nmg_admin_service/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type SportPersist struct {
	app.SportPayload `datastore:"flatten"`
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
		// TODO: Return error.
		return &Context{}
	}
}

func (m *SportPersist) GetAll(q *datastore.Query, mc *Context) *QueryResponse {
	c := []SportPersist{}
	_, err := q.GetAll(mc.GaeCtx, &c)

	ms := make(app.SportMediaCollection, len(c))

	for i, model := range c {
		tm := &app.SportMedia{}
		copier.Copy(tm, model)
		ms[i] = tm
	}

	return &QueryResponse{
		Models: &ms,
		Err:    err,
	}
}

func (m *SportPersist) GetOne(q *datastore.Query, mc *Context) *QueryResponse {
	var tp *SportPersist
	var k *datastore.Key
	var err error

	q = q.KeysOnly()
	for iter := q.Run(mc.GaeCtx); ; {
		var x SportPersist

		key, err := iter.Next(&x)

		if err == datastore.Done {
			break
		}

		tp = &SportPersist{}
		err = datastore.Get(mc.GaeCtx, key, tp)
		k = key
	}

	return &QueryResponse{
		Key:    k,
		Models: tp,
		Err:    err,
	}
}

func (m *SportPersist) Add(ctx *Context) error {
	return m.Set(ctx, datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil))
}

func (m *SportPersist) Set(ctx *Context, key *datastore.Key) error {
	tp := &app.SportPayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return err
}
