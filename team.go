package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
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

	id, err := models.Create(constants.DB_TEAM, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "create")
	}

	return ctx.OKTiny(&app.TeamMediaTiny{id})

	// TeamController_Create: end_implement
}

// Delete runs the delete action.
func (c *TeamController) Delete(ctx *app.DeleteTeamContext) error {
	// TeamController_Delete: start_implement

	if err := models.Delete(constants.DB_TEAM, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "delete")
	}

	return ctx.NoContent()

	// TeamController_Delete: end_implement
}

// List runs the list action.
func (c *TeamController) List(ctx *app.ListTeamContext) error {
	// TeamController_List: start_implement

	coll, err := models.List(constants.DB_TEAM, ctx)

	if err != nil {
		return goa.ErrInternal(err, "endpoint", "list")
	}

	return ctx.OK(coll.(app.TeamMediaCollection))

	// TeamController_List: end_implement
}

// Show runs the show action.
func (c *TeamController) Show(ctx *app.ShowTeamContext) error {
	// TeamController_Show: start_implement

	model, err := models.Read(constants.DB_TEAM, ctx)

	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "show")
	}

	return ctx.OK(model.(*app.TeamMedia))

	// TeamController_Show: end_implement
}

// Update runs the update action.
func (c *TeamController) Update(ctx *app.UpdateTeamContext) error {
	// TeamController_Update: start_implement

	if err := models.Update(constants.DB_TEAM, ctx); err != nil {
		return goa.ErrInternal(err, "endpoint", "update")
	}

	return ctx.NoContent()

	// TeamController_Update: end_implement
}
