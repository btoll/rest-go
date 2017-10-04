package main

import (
	"encoding/json"

	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/goadesign/goa"
)

// SportController implements the Sport resource.
type SportController struct {
	*goa.Controller
}

// NewSportController creates a Sport controller.
func NewSportController(service *goa.Service) *SportController {
	return &SportController{Controller: service.NewController("SportController")}
}

// Create runs the create action.
func (c *SportController) Create(ctx *app.CreateSportContext) error {
	// SportController_Create: start_implement

	id, err := models.GetCtx("SportPersist", ctx).Create()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OKTiny(&app.SportMediaTiny{id})

	// SportController_Create: end_implement
}

// Delete runs the delete action.
func (c *SportController) Delete(ctx *app.DeleteSportContext) error {
	// SportController_Delete: start_implement

	if err := models.GetCtx("SportPersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// SportController_Delete: end_implement
}

// List runs the list action.
func (c *SportController) List(ctx *app.ListSportContext) error {
	// SportController_List: start_implement

	store, err := models.GetCtx("SportPersist", ctx).List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	b, err := json.Marshal(store)

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(b)

	// SportController_List: end_implement
}

// Show runs the show action.
func (c *SportController) Show(ctx *app.ShowSportContext) error {
	// SportController_Show: start_implement

	model, err := models.GetCtx("SportPersist", ctx).Read()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(model.(*app.SportMedia))

	// SportController_Show: end_implement
}

// Update runs the update action.
func (c *SportController) Update(ctx *app.UpdateSportContext) error {
	// SportController_Update: start_implement

	if err := models.GetCtx("SportPersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// SportController_Update: end_implement
}
