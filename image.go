package main

import (
	"fmt"

	"github.com/btoll/rest-go/app"
	"github.com/goadesign/goa"
)

// ImageController implements the Image resource.
type ImageController struct {
	*goa.Controller
}

// NewImageController creates a Image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{Controller: service.NewController("ImageController")}
}

func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	// ImageController_Upload: start_implement

	// ImageController_Upload: end_implement
	return nil
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// ImageController_Upload: start_implement

	// Put your logic here

	fmt.Println()
	fmt.Println("ctx.ID", ctx.ID)
	fmt.Println()

	// ImageController_Upload: end_implement
	res := app.ImageMediaCollection{}
	return ctx.OK(res)
}
