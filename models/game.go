package models

import (
	"context"

	"github.com/dgaedcke/nmg_admin_service/app"
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
		// TODO: Return error.
		return &Context{}
	}
}

func (m *GamePersist) GetAll(q *datastore.Query, mc *Context) *QueryResponse {
	c := []GamePersist{}
	_, err := q.GetAll(mc.GaeCtx, &c)

	ms := make(app.GameMediaCollection, len(c))

	for i, model := range c {
		tm := &app.GameMedia{}
		copier.Copy(tm, model)
		ms[i] = tm
	}

	return &QueryResponse{
		Models: &ms,
		Err:    err,
	}
}

func (m *GamePersist) GetOne(q *datastore.Query, mc *Context) *QueryResponse {
	var tp *GamePersist
	var k *datastore.Key
	var err error

	q = q.KeysOnly()
	for iter := q.Run(mc.GaeCtx); ; {
		var x GamePersist

		key, err := iter.Next(&x)

		if err == datastore.Done {
			break
		}

		tp = &GamePersist{}
		err = datastore.Get(mc.GaeCtx, key, tp)
		k = key
	}

	return &QueryResponse{
		Key:    k,
		Models: tp,
		Err:    err,
	}
}

func (m *GamePersist) Add(ctx *Context) error {
	return m.Set(ctx, datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil))
}

func (m *GamePersist) Set(ctx *Context, key *datastore.Key) error {
	tp := &app.GamePayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return err
}
