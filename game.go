package main

import (
	"fmt"

	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
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

	return nil

	// GameController_Create: end_implement
}

// Delete runs the delete action.
func (c *GameController) Delete(ctx *app.DeleteGameContext) error {
	// GameController_Delete: start_implement

	if err := models.GetCtx("GamePersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// GameController_Delete: end_implement
}

// List runs the list action.
func (c *GameController) List(ctx *app.ListGameContext) error {
	// GameController_List: start_implement

	return nil
	// GameController_List: end_implement
}

// Show runs the show action.
func (c *GameController) Show(ctx *app.ShowGameContext) error {
	// GameController_Show: start_implement

	modelctx := models.GetCtx("GamePersist", ctx)
	rec, err := modelctx.Read()

	fmt.Println()
	fmt.Println("rec", rec)
	fmt.Println()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	//	res := &app.GameMedia{}
	//	copier.Copy(res, modelctx.Persist)
	//
	//	return ctx.OK(res)
	return nil

	// GameController_Show: end_implement
}

// Update runs the update action.
func (c *GameController) Update(ctx *app.UpdateGameContext) error {
	// GameController_Update: start_implement

	if err := models.GetCtx("GamePersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// GameController_Update: end_implement
}
