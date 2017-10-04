package models

import (
	"context"
	"time"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/jinzhu/copier"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type TeamOpeningConfigPersist struct {
	BuyIncrementPrice  float64   `datastore:"buyIncrementPrice,noindex" json:"buyIncrementPrice,omitempty"`
	BuyIncrementQuan   int       `datastore:"buyIncrementQuan,noindex" json:"buyIncrementQuan,omitempty"`
	LiquidationFee     float64   `datastore:"liquidationFee,noindex" json:"liquidationFee,omitempty"`
	OpeningPrice       float64   `datastore:"openingPrice,noindex" json:"openingPrice,omitempty"`
	OpeningShares      int       `datastore:"openingShares,noindex" json:"openingShares,omitempty"`
	SellDecrementPrice float64   `datastore:"sellDecrementPrice,noindex" json:"sellDecrementPrice,omitempty"`
	SellDecrementQuan  int       `datastore:"sellDecrementQuan,noindex" json:"sellDecrementQuan,omitempty"`
	StartTradeDtTm     time.Time `datastore:"startTradeDtTm,noindex" json:"startTradeDtTm,omitempty"`

	Context
}

func (m *TeamOpeningConfigPersist) GetCtx(ctx context.Context) *Context {
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

func (m *TeamOpeningConfigPersist) GetModel() interface{} {
	return &app.TeamOpeningConfigMedia{}
}

func (m *TeamOpeningConfigPersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.TeamOpeningConfigMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *TeamOpeningConfigPersist) SetModel(ctx *Context, key *datastore.Key) error {
	rec := &app.TeamOpeningConfigMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, key, rec)

	return err
}
