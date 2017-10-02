package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/copier"
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

	modelctx := models.GetCtx("SportPersist", ctx)

	if err := modelctx.Create(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.OKTiny(&app.SportMediaTiny{modelctx.ID})
	}

	// SportController_Create: end_implement
}

// Delete runs the delete action.
func (c *SportController) Delete(ctx *app.DeleteSportContext) error {
	// SportController_Delete: start_implement

	if err := models.GetCtx("SportPersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// SportController_Delete: end_implement
}

// List runs the list action.
func (c *SportController) List(ctx *app.ListSportContext) error {
	// SportController_List: start_implement

	modelctx := models.GetCtx("SportPersist", ctx)
	err := modelctx.List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(*modelctx.Persist.(*app.SportMediaCollection))

	// SportController_List: end_implement
	res := app.SportMediaCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *SportController) Show(ctx *app.ShowSportContext) error {
	// SportController_Show: start_implement

	modelctx := models.GetCtx("SportPersist", ctx)
	err := modelctx.Read()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	res := &app.SportMedia{}
	copier.Copy(res, modelctx.Persist)

	return ctx.OK(res)

	// SportController_Show: end_implement
}

// Update runs the update action.
func (c *SportController) Update(ctx *app.UpdateSportContext) error {
	// SportController_Update: start_implement

	if err := models.GetCtx("SportPersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// SportController_Update: end_implement
}
