package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

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

	// Requests can be sent from different sources. If sent from a browser, it will contain
	// a Referer header, and we must convert the base64 encoded image to binary.
	if ctx.Referer() == "" {
		reader, err := ctx.MultipartReader()
		if err != nil {
			return goa.ErrBadRequest(err, "endpoint", "upload")
		}

		for {
			part, err := reader.NextPart()

			if err != io.EOF {
				storageobject := client.Bucket(bucketname).Object(fmt.Sprintf("%s/%s/%s", ctx.Entity, ctx.ID, part.FileName()))
				writer := storageobject.NewWriter(gaeCtx)
				writer.ContentType = "image/*"

				//writer.Metadata = map[string]string{
				//    "x-googl-acl": "authenticated-read, public-read",
				//}

				_, err = io.Copy(writer, part)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}

				err = writer.Close()
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}

				acl := storageobject.ACL()
				if err = acl.Set(gaeCtx, storage.AllUsers /*storage.AllAuthenticatedUsers*/, storage.RoleReader); err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
			} else {
				break
			}
		}
	} else {
		type fileobject struct {
			Filename, Contents string
		}

		// https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream
		dec := json.NewDecoder(ctx.Body)
		// read open bracket
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v\n", t, t)

		// while the array contains values
		for dec.More() {
			var m fileobject
			// decode an array value (Message)
			err := dec.Decode(&m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%v: %v\n", m.Filename, m.Contents)
		}

		// read closing bracket
		t, err = dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		fmt.Printf("%T: %v\n", t, t)
		fmt.Println()
	}

	return nil
	// ImageController_Upload: end_implement
}
