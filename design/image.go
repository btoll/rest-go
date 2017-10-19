package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Image", func() {
	BasePath("/image")

	Action("show", func() {
		Routing(GET("/:entity/:id"))
		Description("Show all images for a particular team")
		Params(func() {
			Param("entity", String, "event|sport|team")
			Param("id", String, "Entity ID")
		})
		Response(OK, ImageMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("list", func() {
		//        Routing(GET("/list/:entity"))
		Routing(GET("/:entity"))
		Params(func() {
			Param("entity", String, "event|sport|team")
		})
		Description("Get all images of all teams")
		//		Response(OK, "application/json")
		Response(OK, CollectionOf(ImageMedia))
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("upload", func() {
		Routing(PUT("/:entity/:id"))
		Description("Upload multiple images in multipart request")
		Params(func() {
			Param("entity", String, "event|sport|team")
			Param("id", String, "Entity ID")
		})
		Response(OK, CollectionOf(ImageMedia))
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
