package main

import (
	"fmt"

	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
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

	return nil

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

	return nil

	// TeamOpeningConfigController_List: end_implement
}

// Show runs the show action.
func (c *TeamOpeningConfigController) Show(ctx *app.ShowTeamOpeningConfigContext) error {
	// TeamOpeningConfigController_Show: start_implement

	modelctx := models.GetCtx("TeamOpeningConfigPersist", ctx)
	rec, err := modelctx.Read()

	fmt.Println()
	fmt.Println("rec", rec)
	fmt.Println()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	//	res := &app.TeamOpeningConfigMedia{}
	//	copier.Copy(res, modelctx.Persist)
	//
	//	return ctx.OK(res)

	return nil
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
