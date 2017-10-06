package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
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

	id, err := models.GetCtx(constants.DB_TEAM_OPENING_CONFIG, ctx).Create()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OKTiny(&app.TeamOpeningConfigMediaTiny{id})

	// TeamOpeningConfigController_Create: end_implement
}

// Delete runs the delete action.
func (c *TeamOpeningConfigController) Delete(ctx *app.DeleteTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Delete: start_implement

	if err := models.GetCtx(constants.DB_TEAM_OPENING_CONFIG, ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// TeamOpeningConfigController_Delete: end_implement
}

// List runs the list action.
func (c *TeamOpeningConfigController) List(ctx *app.ListTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_List: start_implement

	b, err := models.GetCtx(constants.DB_TEAM_OPENING_CONFIG, ctx).List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(b)

	// TeamOpeningConfigController_List: end_implement
}

// Show runs the show action.
func (c *TeamOpeningConfigController) Show(ctx *app.ShowTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Show: start_implement

	model, err := models.GetCtx(constants.DB_TEAM_OPENING_CONFIG, ctx).Read()

	if err != nil {
		return ctx.BadRequest(err)
	}

	return ctx.OK(model.(*app.TeamOpeningConfigMedia))

	// TeamOpeningConfigController_Show: end_implement
}

// Update runs the update action.
func (c *TeamOpeningConfigController) Update(ctx *app.UpdateTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Update: start_implement

	if err := models.GetCtx(constants.DB_TEAM_OPENING_CONFIG, ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// TeamOpeningConfigController_Update: end_implement
}
