package models

import (
	"context"

	"github.com/btoll/rest-go/app"
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

func (m *EventPersist) GetModelCollection(ctx *Context) (interface{}, error) {
	return nil, nil
}

func (m *EventPersist) GetModel() interface{} {
	return &app.EventMedia{}
}

func (m *EventPersist) Set(ctx *Context, key *datastore.Key) (string, error) {
	rec := &app.EventMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return "", err
}
