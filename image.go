package main

import (
	"io"

	"cloud.google.com/go/storage"
	"github.com/btoll/rest-go/app"
	"github.com/goadesign/goa"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
)

// ImageController implements the Image resource.
type ImageController struct {
	*goa.Controller
}

// NewImageController creates a Image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{Controller: service.NewController("ImageController")}
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// ImageController_Upload: start_implement

	/* ----------------------------------------------------------- */
	gaeCtx := appengine.NewContext(ctx.Request)

	client, err := storage.NewClient(gaeCtx)
	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "upload")
	}
	defer client.Close()

	bucketname, err := file.DefaultBucketName(gaeCtx)
	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "upload")
	}
	/* ----------------------------------------------------------- */

	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest(err, "endpoint", "upload")
	}

	for {
		part, err := reader.NextPart()

		if err != io.EOF {
			writer := client.Bucket(bucketname).Object(ctx.Entity + "/" + ctx.ID + "/" + part.FileName()).NewWriter(gaeCtx)

			_, err = io.Copy(writer, part)
			if err != nil {
				return goa.ErrBadRequest(err, "endpoint", "upload")
			}

			err = writer.Close()
			if err != nil {
				return goa.ErrBadRequest(err, "endpoint", "upload")
			}
		} else {
			break
		}
	}

	return nil
	// ImageController_Upload: end_implement
}