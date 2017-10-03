package models

import (
	"context"

	"github.com/btoll/rest-go/app"
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

func (m *GamePersist) GetModelCollection(ctx *Context) (interface{}, error) {
	return nil, nil
}

func (m *GamePersist) GetModel() interface{} {
	//	var tp *GamePersist
	//	var k *datastore.Key
	//	var err error
	//
	//	q = q.KeysOnly()
	//	for iter := q.Run(mc.GaeCtx); ; {
	//		var x GamePersist
	//
	//		key, err := iter.Next(&x)
	//
	//		if err == datastore.Done {
	//			break
	//		}
	//
	//		tp = &GamePersist{}
	//		err = datastore.Get(mc.GaeCtx, key, tp)
	//		k = key
	//	}
	//
	//	return &QueryResponse{
	//		Key:    k,
	//		Models: tp,
	//		Err:    err,
	//	}
	return &app.GameMedia{}
}

func (m *GamePersist) Set(ctx *Context, key *datastore.Key) (string, error) {
	rec := &app.GameMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return "", err
}
