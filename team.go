package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/copier"
)

// TeamController implements the Team resource.
type TeamController struct {
	*goa.Controller
}

// NewTeamController creates a Team controller.
func NewTeamController(service *goa.Service) *TeamController {
	return &TeamController{Controller: service.NewController("TeamController")}
}

// Create runs the create action.
func (c *TeamController) Create(ctx *app.CreateTeamContext) error {
	// TeamController_Create: start_implement

	modelctx := models.GetCtx("TeamPersist", ctx)

	if err := modelctx.Create(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.OKTiny(&app.TeamMediaTiny{modelctx.ID})
	}

	// TeamController_Create: end_implement
}

// Delete runs the delete action.
func (c *TeamController) Delete(ctx *app.DeleteTeamContext) error {
	// TeamController_Delete: start_implement

	if err := models.GetCtx("TeamPersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// TeamController_Delete: end_implement
}

// List runs the list action.
func (c *TeamController) List(ctx *app.ListTeamContext) error {
	// TeamController_List: start_implement

	modelctx := models.GetCtx("TeamPersist", ctx)
	err := modelctx.List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(*modelctx.Persist.(*app.TeamMediaCollection))

	// TeamController_List: end_implement
}

// Show runs the show action.
func (c *TeamController) Show(ctx *app.ShowTeamContext) error {
	// TeamController_Show: start_implement

	modelctx := models.GetCtx("TeamPersist", ctx)
	err := modelctx.Read()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	res := &app.TeamMedia{}
	copier.Copy(res, modelctx.Persist)

	return ctx.OK(res)

	// TeamController_Show: end_implement
}

// Update runs the update action.
func (c *TeamController) Update(ctx *app.UpdateTeamContext) error {
	// TeamController_Update: start_implement

	if err := models.GetCtx("TeamPersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// TeamController_Update: end_implement
}
