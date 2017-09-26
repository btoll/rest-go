package main

import (
	"github.com/dgaedcke/nmg_admin_service/app"
	"github.com/dgaedcke/nmg_admin_service/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/copier"
)

// TeamOpeningConfigController implements the TeamOpeningConfig resource.
type TeamOpeningConfigController struct {
	*goa.Controller
}

// NewTeamOpeningConfigController creates a TeamOpeningConfig controller.
func NewTeamOpeningConfigController(service *goa.Service) *TeamOpeningConfigController {
	return &TeamOpeningConfigController{Controller: service.NewController("TeamOpeningConfigController")}
}

// Create runs the create action.
func (c *TeamOpeningConfigController) Create(ctx *app.CreateTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Create: start_implement

	modelctx := models.GetCtx("TeamOpeningConfigPersist", ctx)

	if err := modelctx.Create(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.OKTiny(&app.TeamOpeningConfigMediaTiny{modelctx.ID})
	}

	// TeamOpeningConfigController_Create: end_implement
}

// Delete runs the delete action.
func (c *TeamOpeningConfigController) Delete(ctx *app.DeleteTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Delete: start_implement

	if err := models.GetCtx("TeamOpeningConfigPersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// TeamOpeningConfigController_Delete: end_implement
}

// List runs the list action.
func (c *TeamOpeningConfigController) List(ctx *app.ListTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_List: start_implement

	modelctx := models.GetCtx("TeamOpeningConfigPersist", ctx)
	err := modelctx.List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(*modelctx.Persist.(*app.TeamOpeningConfigMediaCollection))

	// TeamOpeningConfigController_List: end_implement
}

// Show runs the show action.
func (c *TeamOpeningConfigController) Show(ctx *app.ShowTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Show: start_implement

	modelctx := models.GetCtx("TeamOpeningConfigPersist", ctx)
	err := modelctx.Read()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	res := &app.TeamOpeningConfigMedia{}
	copier.Copy(res, modelctx.Persist)

	return ctx.OK(res)

	// TeamOpeningConfigController_Show: end_implement
}

// Update runs the update action.
func (c *TeamOpeningConfigController) Update(ctx *app.UpdateTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Update: start_implement

	if err := models.GetCtx("TeamOpeningConfigPersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// TeamOpeningConfigController_Update: end_implement
}
