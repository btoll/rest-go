package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Image", func() {
	BasePath("/image")

	Action("upload", func() {
		Routing(PUT("/:id"))
		Description("Upload multiple images in multipart request")
		Params(func() {
			Param("id", String, "event|sport|team")
		})
		Response(OK, CollectionOf(ImageMedia))
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Show an image's metadata")
		Params(func() {
			Param("id", Integer, "Image ID")
		})
		Response(OK, ImageMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

var ImageMedia = MediaType("application/nmgapi.imageentity", func() {
	Description("Image metadata")
	TypeName("ImageMedia")

	Attributes(func() {
		Attribute("id", Integer, "Image ID")
		Attribute("filename", String, "Image filename")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Required("id", "filename", "uploaded_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("filename")
		Attribute("uploaded_at")
	})
})
