package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"strings"

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

	// https://stackoverflow.com/questions/22945486/golang-converting-image-image-to-byte
	// https://stackoverflow.com/questions/33319759/go-base64-image-decode
	// https://stackoverflow.com/questions/31031589/illegal-base64-data-at-input-byte-4-when-using-base64-stdencoding-decodestrings

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

	bucket := client.Bucket(bucketname)

	// Let's define some lambdas. Since they're closures they will have side effects,
	// but this won't have any negative consequences since there's no concurrency.
	var writer *storage.Writer
	var storageobject *storage.ObjectHandle
	createEssentials := func(s string) {
		storageobject = bucket.Object(fmt.Sprintf("%s/%s/%s", ctx.Entity, ctx.ID, s))
		writer = storageobject.NewWriter(gaeCtx)
		writer.ContentType = "image/*"
	}

	copyAndClose := func(src io.Reader) error {
		_, err = io.Copy(writer, src)
		if err != nil {
			return err
		}
		// Closing the writer will initiate the request to GAE.
		return writer.Close()
	}

	setAcl := func() error {
		return storageobject.ACL().Set(gaeCtx, storage.AllUsers /*storage.AllAuthenticatedUsers*/, storage.RoleReader)
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
				createEssentials(part.FileName())
				err = copyAndClose(part)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
				err = setAcl()
				if err != nil {
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
		_, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		// while the array contains values
		for dec.More() {
			var m fileobject
			// decode an array value (Message)
			err := dec.Decode(&m)
			if err != nil {
				log.Fatal(err)
			}

			// Extract the actual base64-encoded string from the data uri scheme.
			//
			//      data:image/jpeg;base64,{...}
			//      data:image/png;base64,{...}
			//      ...etc...
			//
			base64data := m.Contents[strings.IndexByte(m.Contents, ',')+1:]

			buf := new(bytes.Buffer)
			reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64data))

			// Extract the image type from the data uri scheme.
			//
			//      {data:image/jpeg;base64,}...
			//      {data:image/png;base64,}...
			//      ...etc...
			//
			imageType := m.Contents[5:strings.Index(m.Contents, ";base64,")]

			// TODO: DRY?
			switch imageType {
			case "image/gif":
				img, err := gif.DecodeAll(reader)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
				err = gif.EncodeAll(buf, img)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
			case "image/jpeg":
				img, _, err := image.Decode(reader)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
				err = jpeg.Encode(buf, img, nil)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
			case "image/png":
				img, err := png.Decode(reader)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
				png.Encode(buf, img)
				if err != nil {
					return goa.ErrBadRequest(err, "endpoint", "upload")
				}
			}

			createEssentials(m.Filename)
			err = copyAndClose(bytes.NewReader(buf.Bytes()))
			if err != nil {
				return goa.ErrBadRequest(err, "endpoint", "upload")
			}
			err = setAcl()
			if err != nil {
				return goa.ErrBadRequest(err, "endpoint", "upload")
			}
		}

		// read closing bracket
		_, err = dec.Token()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
	// ImageController_Upload: end_implement
}
