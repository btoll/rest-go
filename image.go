package main

import (
	"bytes"
	"io"

	"cloud.google.com/go/storage"
	"github.com/btoll/rest-go/app"
	"github.com/goadesign/goa"
	"google.golang.org/appengine"
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

	/* ----------------------------------------------------------- */
	gaeCtx := appengine.NewContext(ctx.Request)
	//	bucketname, err := file.DefaultBucketName(gaeCtx)

	client, err := storage.NewClient(gaeCtx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//	wc := client.Bucket(bucketname).Object("derp").NewWriter(gaeCtx)
	wc := client.Bucket("foobar-179819.appspot.com").Object("derp").NewWriter(gaeCtx)
	wc.ContentType = "text/plain"
	/* ----------------------------------------------------------- */

	reader, _ := ctx.MultipartReader()
	pr, pw := io.Pipe()
	defer pr.Close()

	go func() {
		buf := new(bytes.Buffer)

		for {
			part, err := reader.NextPart()

			if err != io.EOF {
				// TODO: Don't want to allocate...
				_, err := buf.ReadFrom(part)
				if err != nil {
					panic(err)
				}

				buf.WriteTo(pw)
			} else {
				pw.Close()
				break
			}
		}
	}()

	//	io.Copy(io.MultiWriter(os.Stdout, os.Stderr, wc), pr)
	io.Copy(wc, pr)
	wc.Close()

	// ImageController_Upload: end_implement
	res := app.ImageMediaCollection{}
	return ctx.OK(res)
}
