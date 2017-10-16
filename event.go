package main

import (
	"github.com/btoll/rest-go/app"
	"github.com/btoll/rest-go/models"
	"github.com/dgaedcke/nmg_shared/constants"
	"github.com/goadesign/goa"
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

	id, err := models.Create(constants.DB_EVENT, ctx)

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OKTiny(&app.EventMediaTiny{id})

	// EventController_Create: end_implement
	res := &app.EventMediaTiny{}
	return ctx.OKTiny(res)
}

// Delete runs the delete action.
func (c *EventController) Delete(ctx *app.DeleteEventContext) error {
	// EventController_Delete: start_implement

	if err := models.Delete(constants.DB_EVENT, ctx); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// EventController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *EventController) List(ctx *app.ListEventContext) error {
	// EventController_List: start_implement

	b, err := models.List(constants.DB_EVENT, ctx)

	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(b)

	// EventController_List: end_implement
	return nil
}

// Show runs the show action.
func (c *EventController) Show(ctx *app.ShowEventContext) error {
	// EventController_Show: start_implement

	model, err := models.Read(constants.DB_EVENT, ctx)

	if err != nil {
		return ctx.BadRequest(err)
	}

	return ctx.OK(model.(*app.EventMedia))

	// EventController_Show: end_implement
	res := &app.EventMedia{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *EventController) Update(ctx *app.UpdateEventContext) error {
	// EventController_Update: start_implement

	if err := models.Update(constants.DB_EVENT, ctx); err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()

	// EventController_Update: end_implement
	return nil
}
