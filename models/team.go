package models

import (
	"context"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamPersist struct {
	// Team Win-Loss Record
	CurrentWinRecord string `datastore:"currentWinRecord,noindex" json:"currentWinRecord,omitempty"`
	// Team Logo
	FullLogo string `datastore:"fullLogo,noindex" json:"fullLogo,omitempty"`
	// Sport HomeTown ID
	HomeTownID string `datastore:"homeTownId,noindex" json:"homeTownId,omitempty"`
	// Team Icon
	IconName string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Team name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Team Nickname
	ShortName string `datastore:"shortName,noindex" json:"shortName,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`

	Context
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

func (m *TeamPersist) GetModel() interface{} {
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

func (m *TeamPersist) SetModel(ctx *Context, key *datastore.Key) error {
	rec := &app.TeamMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return err
}
