package main

import (
	"io"
	"os"

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

// List runs the list action.
func (c *ImageController) List(ctx *app.ListImageContext) error {
	// ImageController_List: start_implement

	// Put your logic here

	// ImageController_List: end_implement
	res := app.ImageMediaCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	// ImageController_Show: start_implement

	// Put your logic here

	// ImageController_Show: end_implement
	res := &app.ImageMedia{}
	return ctx.OK(res)
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// ImageController_Upload: start_implement

	reader, _ := ctx.MultipartReader()
	pr, pw := io.Pipe()
	defer pr.Close()

	go func() {
		// TODO: Don't want to allocate.
		b := make([]byte, 4096)

		for {
			part, err := reader.NextPart()

			if err != io.EOF {
				part.Read(b)
				pw.Write(b)
			} else {
				pw.Close()
				break
			}

			/*
				if part.FormName() == "delete" {
					buf := new(bytes.Buffer)
					buf.ReadFrom(part)
					log.Println("delete is: ", buf.String())
				}
			*/
		}
	}()

	io.Copy(os.Stdout, pr)

	// ImageController_Upload: end_implement
	res := app.ImageMediaCollection{}
	return ctx.OK(res)
}
