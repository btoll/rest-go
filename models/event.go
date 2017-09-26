package models

import (
	"context"

	"github.com/dgaedcke/nmg_admin_service/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type EventPersist struct {
	app.EventPayload
	Context
}

func (m *EventPersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateEventContext:
		return &Context{
			Kind:    "EventPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteEventContext:
		return &Context{
			Kind:   "EventPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListEventContext:
		return &Context{
			Kind:   "EventPersist",
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowEventContext:
		return &Context{
			Kind:   "EventPersist",
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateEventContext:
		return &Context{
			Kind:    "EventPersist",
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error.
		return &Context{}
	}
}

func (m *EventPersist) GetAll(q *datastore.Query, mc *Context) *QueryResponse {
	c := []EventPersist{}
	_, err := q.GetAll(mc.GaeCtx, &c)

	ms := make(app.EventMediaCollection, len(c))

	for i, model := range c {
		tm := &app.EventMedia{}
		copier.Copy(tm, model)
		ms[i] = tm
	}

	return &QueryResponse{
		Models: &ms,
		Err:    err,
	}
}

func (m *EventPersist) GetOne(q *datastore.Query, mc *Context) *QueryResponse {
	var tp *EventPersist
	var k *datastore.Key
	var err error

	q = q.KeysOnly()
	for iter := q.Run(mc.GaeCtx); ; {
		var x EventPersist

		key, err := iter.Next(&x)

		if err == datastore.Done {
			break
		}

		tp = &EventPersist{}
		err = datastore.Get(mc.GaeCtx, key, tp)
		k = key
	}

	return &QueryResponse{
		Key:    k,
		Models: tp,
		Err:    err,
	}
}

func (m *EventPersist) Add(ctx *Context) error {
	return m.Set(ctx, datastore.NewKey(ctx.GaeCtx, ctx.Kind, ctx.ID, 0, nil))
}

func (m *EventPersist) Set(ctx *Context, key *datastore.Key) error {
	tp := &app.EventPayload{}

	copier.Copy(tp, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, tp)

	return err
}
