package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Sport", func() {
	BasePath("/sport")
	// Seems that goa doesn't like setting DefaultMedia here at the top-level when the MediaType has multiple Views.
	//	DefaultMedia(SportMedia)
	Description("Describes a sport.")

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new sport.")
		Payload(SportPayload)
		Response(OK, func() {
			Status(200)
			Media(SportMedia, "tiny")
		})
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Sport ID")
		})
		Description("Get a sport by id.")
		Response(OK, SportMedia)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(SportPayload)
		Params(func() {
			Param("id", String, "Sport ID")
		})
		Description("Update a sport by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Sport ID")
		})
		Description("Delete a sport by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all sports")
		Response(OK, CollectionOf(SportMedia))
	})
})

var SportPayload = Type("SportPayload", func() {
	Description("Sport Description.")

	Attribute("id", String, "ID", func() {
		Metadata("struct:tag:datastore", "id,noindex")
		Metadata("struct:tag:json", "id,omitempty")
	})
	Attribute("name", String, "Sport name", func() {
		Metadata("struct:tag:datastore", "name,noindex")
		Metadata("struct:tag:json", "name,omitempty")
	})
	Attribute("active", Boolean, "Is in season?", func() {
		Metadata("struct:tag:datastore", "active,noindex")
		Metadata("struct:tag:json", "active")
	})
	Attribute("maxPreSplitPrice", Number, "", func() {
		Metadata("struct:tag:datastore", "maxPreSplitPrice,noindex")
		Metadata("struct:tag:json", "maxPreSplitPrice")
	})
	Attribute("gameTerm", String, "Game", func() {
		Metadata("struct:tag:datastore", "gameTerm,noindex")
		Metadata("struct:tag:json", "gameTerm,omitempty")
	})
	Attribute("eventTerm", String, "Tournament", func() {
		Metadata("struct:tag:datastore", "eventTerm,noindex")
		Metadata("struct:tag:json", "eventTerm,omitempty")
	})

	Required("name", "active", "maxPreSplitPrice", "gameTerm", "eventTerm")
})

var SportMedia = MediaType("application/nmgapi.sportentity", func() {
	Description("Sport response")
	TypeName("SportMedia")
	ContentType("application/json")
	Reference(SportPayload)

	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("active")
		Attribute("maxPreSplitPrice")
		Attribute("gameTerm")
		Attribute("eventTerm")

		Required("id", "name", "active", "maxPreSplitPrice", "gameTerm", "eventTerm")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("active")
		Attribute("maxPreSplitPrice")
		Attribute("gameTerm")
		Attribute("eventTerm")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new sports.")
		Attribute("name")
	})
})
