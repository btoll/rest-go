package models

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/appengine/datastore"
)

type App interface {
	Create() error
	Read() (Model, error)
	Update() error
	Delete() error
}

type Model interface {
	GetCtx(ctx context.Context) *Context
	GetAll(q *datastore.Query, mc *Context) *QueryResponse
	GetOne(q *datastore.Query, mc *Context) *QueryResponse
	Add(ctx *Context) error
	Set(ctx *Context, key *datastore.Key) error
}

type Context struct {
	GaeCtx  context.Context
	Payload interface{}
	Persist interface{}
	Kind    string
	ID      string
	App
}

type QueryResponse struct {
	Key    *datastore.Key
	Models interface{}
	Err    error
}

const (
	GAME                = "GamePersist"
	EVENT               = "EventPersist"
	SPORT               = "SportPersist"
	TEAM                = "TeamPersist"
	TEAM_OPENING_CONFIG = "TeamOpeningConfigPersist"
)

func ModelFactory(m string) Model {
	switch m {
	case GAME:
		return new(GamePersist)
	case EVENT:
		return new(EventPersist)
	case SPORT:
		return new(SportPersist)
	case TEAM:
		return new(TeamPersist)
	case TEAM_OPENING_CONFIG:
		return new(TeamOpeningConfigPersist)
	}

	// TODO: Return error?
	return nil
}

func GetCtx(kind string, ctx context.Context) *Context {
	return ModelFactory(kind).GetCtx(ctx)
}

/* -----------------------------------------------------------
    CRUD(L)
----------------------------------------------------------- */

func (mc *Context) Create() error {
	mc.ID = GetRandId()
	return ModelFactory(mc.Kind).Add(mc)
}

func (mc *Context) Read() error {
	res := mc.GetOne()

	if res.Err != nil {
		return res.Err
	}

	mc.Persist = res.Models
	return nil
}

func (mc *Context) Update() error {
	res := mc.GetOne()

	if res.Err != nil {
		return res.Err
	}

	return ModelFactory(mc.Kind).Set(mc, res.Key)
}

func (mc *Context) Delete() error {
	res := mc.GetOne()

	if res.Err != nil {
		return res.Err
	}

	return datastore.Delete(mc.GaeCtx, res.Key)
}

func (mc *Context) List() error {
	res := mc.GetAll()

	mc.Persist = res.Models
	return res.Err
}

/* -----------------------------------------------------------
    ERROR (TODO)
----------------------------------------------------------- */

/* -----------------------------------------------------------
    CRUD Abstractions
----------------------------------------------------------- */

// TODO: FP!
func (mc *Context) GetOne() *QueryResponse {
	q := datastore.NewQuery(mc.Kind).Filter("id =", mc.ID)
	return ModelFactory(mc.Kind).GetOne(q, mc)
}

func (mc *Context) GetAll() *QueryResponse {
	q := datastore.NewQuery(mc.Kind)
	return ModelFactory(mc.Kind).GetAll(q, mc)
}

/* -----------------------------------------------------------
    MISC
----------------------------------------------------------- */

// Temporary to mock the id.
func GetRandId() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Intn(100000))
}
