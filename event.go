package main

import (
	"github.com/dgaedcke/nmg_admin_service/app"
	"github.com/dgaedcke/nmg_admin_service/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/copier"
)

// EventController implements the Event resource.
type EventController struct {
	*goa.Controller
}

// NewEventController creates a Event controller.
func NewEventController(service *goa.Service) *EventController {
	return &EventController{Controller: service.NewController("EventController")}
}

// Create runs the create action.
func (c *EventController) Create(ctx *app.CreateEventContext) error {
	// EventController_Create: start_implement

	modelctx := models.GetCtx("EventPersist", ctx)

	if err := modelctx.Create(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.OKTiny(&app.EventMediaTiny{modelctx.ID})
	}

	// EventController_Create: end_implement
}

// Delete runs the delete action.
func (c *EventController) Delete(ctx *app.DeleteEventContext) error {
	// EventController_Delete: start_implement

	if err := models.GetCtx("EventPersist", ctx).Delete(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// EventController_Delete: end_implement
}

// List runs the list action.
func (c *EventController) List(ctx *app.ListEventContext) error {
	// EventController_List: start_implement

	modelctx := models.GetCtx("EventPersist", ctx)
	err := modelctx.List()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(*modelctx.Persist.(*app.EventMediaCollection))

	// EventController_List: end_implement
}

// Show runs the show action.
func (c *EventController) Show(ctx *app.ShowEventContext) error {
	// EventController_Show: start_implement

	modelctx := models.GetCtx("EventPersist", ctx)
	err := modelctx.Read()

	if err != nil {
		return ctx.InternalServerError(err)
	}

	// TODO: Hacky?
	res := &app.EventMedia{}
	copier.Copy(res, modelctx.Persist)

	return ctx.OK(res)

	// EventController_Show: end_implement
}

// Update runs the update action.
func (c *EventController) Update(ctx *app.UpdateEventContext) error {
	// EventController_Update: start_implement

	if err := models.GetCtx("EventPersist", ctx).Update(); err != nil {
		return ctx.InternalServerError(err)
	} else {
		return ctx.NoContent()
	}

	// EventController_Update: end_implement
}
