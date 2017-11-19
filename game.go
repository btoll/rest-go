package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
)

// GameController implements the Game resource.
type GameController struct {
	*goa.Controller
}

// NewGameController creates a Game controller.
func NewGameController(service *goa.Service) *GameController {
	return &GameController{Controller: service.NewController("GameController")}
}

// Create runs the create action.
func (c *GameController) Create(ctx *app.CreateGameContext) error {
	// GameController_Create: start_implement

	id, err := models.Create(constants.DB_GAME, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "create")
	}

	return ctx.OKTiny(&app.GameMediaTiny{id})

	// GameController_Create: end_implement
}

// Delete runs the delete action.
func (c *GameController) Delete(ctx *app.DeleteGameContext) error {
	// GameController_Delete: start_implement

	if err := models.Delete(constants.DB_GAME, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "delete")
	}

	return ctx.NoContent()

	// GameController_Delete: end_implement
}

// List runs the list action.
func (c *GameController) List(ctx *app.ListGameContext) error {
	// GameController_List: start_implement

	coll, err := models.List(constants.DB_GAME, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "list")
	}

	return ctx.OK(coll.(app.GameMediaCollection))

	// GameController_List: end_implement
}

// Show runs the show action.
func (c *GameController) Show(ctx *app.ShowGameContext) error {
	// GameController_Show: start_implement

	model, err := models.Read(constants.DB_GAME, ctx)

	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "show")
	}

	return ctx.OK(model.(*app.GameMedia))

	// GameController_Show: end_implement
}

// Update runs the update action.
func (c *GameController) Update(ctx *app.UpdateGameContext) error {
	// GameController_Update: start_implement

	if err := models.Update(constants.DB_GAME, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "update")
	}

	return ctx.NoContent()

	// GameController_Update: end_implement
}
