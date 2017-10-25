package main

import (
	"encoding/json"

	"github.com/btoll/rest-go/app"
	"github.com/dgaedcke/nmg_shared/enum"
	"github.com/goadesign/goa"
)

// EnumController implements the Enum resource.
type EnumController struct {
	*goa.Controller
}

// NewEnumController creates a Enum controller.
func NewEnumController(service *goa.Service) *EnumController {
	return &EnumController{Controller: service.NewController("EnumController")}
}

// List runs the list action.
func (c *EnumController) List(ctx *app.ListEnumContext) error {
	// EnumController_List: start_implement

	b, err := json.Marshal(enum.GetAll())
	if err != nil {
		return goa.ErrInternal(err, "endpoint", "list")
	}

	return ctx.OK(b)

	// EnumController_List: end_implement
}
