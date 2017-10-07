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

type GamePersist struct {
	// Event ID
	EventID *string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Any public id for this game, not unique
	ExternalID string `datastore:"externalId,noindex" json:"externalId,omitempty"`
	// Favorite team id
	FavTeamID      string    `datastore:"favTeamId,noindex" json:"favTeamId,omitempty"`
	FinishedAtDtTm time.Time `datastore:"finishedAtDtTm,noindex" json:"finishedAtDtTm,omitempty"`
	// TeamGameStatus.preGame || tradeable || gameOn || ended
	GameStatus string `datastore:"gameStatus,noindex" json:"gameStatus,omitempty"`
	// Name of location
	Location string `datastore:"location,noindex" json:"location,omitempty"`
	// True GPS location
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// TeamGameStatus.eliminated
	LoserProgressStatus string `datastore:"loserProgressStatus,noindex" json:"loserProgressStatus,omitempty"`
	// 0
	OddsForFav float64   `datastore:"oddsForFav,noindex" json:"oddsForFav,omitempty"`
	PlayDtTm   time.Time `datastore:"playDtTm,noindex" json:"playDtTm,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
	// Public free form name
	Title       string `datastore:"title,noindex" json:"title,omitempty"`
	UnderTeamID string `datastore:"underTeamId,noindex" json:"underTeamId,omitempty"`
	// Empty until game completed
	WinnerTeamID string `datastore:"winnerTeamId,noindex" json:"winnerTeamId,omitempty"`

	*Context
}

func (m *GamePersist) AllocateID(ctx *Context) error {
	return Allocate(ctx)
}

func (m *GamePersist) GetCtx(ctx context.Context) *Context {
	switch t := ctx.(type) {
	case *app.CreateGameContext:
		return &Context{
			Kind:    constants.DB_GAME,
			GaeCtx:  appengine.NewContext(t.Request),
			Payload: t.Payload,
		}
	case *app.DeleteGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.ListGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
		}
	case *app.ShowGameContext:
		return &Context{
			Kind:   constants.DB_GAME,
			GaeCtx: appengine.NewContext(t.Request),
			ID:     t.ID,
		}
	case *app.UpdateGameContext:
		return &Context{
			Kind:    constants.DB_GAME,
			GaeCtx:  appengine.NewContext(t.Request),
			ID:      t.ID,
			Payload: t.Payload,
		}
	default:
		// TODO: Return error?
		return &Context{}
	}
}

func (m *GamePersist) GetModelInstance() interface{} {
	return &app.GameMedia{}
}

func (m *GamePersist) GetModelCollection(ctx *Context) ([]*datastore.Key, interface{}, error) {
	c := []app.GameMedia{}
	keys, err := datastore.NewQuery(ctx.Kind).GetAll(ctx.GaeCtx, &c)

	if err != nil {
		return nil, nil, err
	}

	return keys, c, nil
}

func (m *GamePersist) SetModel(ctx *Context) error {
	rec := &app.GameMedia{}

	copier.Copy(rec, ctx.Payload)
	_, err := datastore.Put(ctx.GaeCtx, GetKey(ctx), rec)

	return err
}
