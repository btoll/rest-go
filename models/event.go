package models

import (
	"context"
	"time"

	"github.com/btoll/rest-go/app"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type EventPersist struct {
	EndDtTm time.Time `datastore:"endDtTm,noindex" json:"endDtTm,omitempty"`
	// Not guaranteed to be unique
	EventID string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Location.defaultLoc.id
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// e.g., March Madness
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Sport ID
	SportID   string    `datastore:"sportId,noindex" json:"sportId,omitempty"`
	StartDtTm time.Time `datastore:"startDtTm,noindex" json:"startDtTm,omitempty"`
	SubTitle  string    `datastore:"subTitle,noindex" json:"subTitle,omitempty"`
	// EventAdvanceMethod.singleElimination || doubleElim || bestOf
	TeamAdvanceMethod string `datastore:"teamAdvanceMethod,noindex" json:"teamAdvanceMethod,omitempty"`

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
		// TODO: Return error?
		return &Context{}
	}
}

func (m *EventPersist) GetModel() interface{} {
	return &app.EventMedia{}
}

func (m *EventPersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.EventMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *EventPersist) SetModel(ctx *Context, key *datastore.Key) error {
	rec := &app.EventMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return err
}
