package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
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

			storageobject := client.Bucket(bucketname).Object(fmt.Sprintf("%s/%s/%s", ctx.Entity, ctx.ID, m.Filename))
			writer := storageobject.NewWriter(gaeCtx)
			writer.ContentType = "image/*"

			// Extract the image type from the data uri scheme.
			//
			//      {data:image/jpeg;base64,}...
			//      {data:image/png;base64,}...
			//      ...etc...
			//
			imageType := m.Contents[5:strings.Index(m.Contents, ";base64,")]

			// Extract the actual base64-encoded string from the data uri scheme.
			//
			//      data:image/jpeg;base64,{...}
			//      data:image/png;base64,{...}
			//      ...etc...
			//
			base64data := m.Contents[strings.IndexByte(m.Contents, ',')+1:]

			buf := new(bytes.Buffer)

			switch imageType {
			//			case "image/gif":
			//				image, err := gif.Decode(m.Contents)

			case "image/jpeg":
				reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64data))
				img, _, err := image.Decode(reader)
				if err != nil {
					fmt.Println(err)
				}

				jpeg.Encode(buf, img, nil)

				//			case "image/png":
				//				image, err := png.Decode(m.Contents)
			}

			_, err = io.Copy(writer, bytes.NewReader(buf.Bytes()))
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
