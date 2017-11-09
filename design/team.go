package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Team", func() {
	BasePath("/team")
	// Seems that goa doesn't like setting DefaultMedia here at the top-level when the MediaType has multiple Views.
	//	DefaultMedia(TeamMedia)
	Description("Describes a sport team.")

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new sports team.")
		Payload(TeamPayload)
		Response(OK, func() {
			Status(200)
			Media(TeamMedia, "tiny")
		})
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Team ID")
		})
		Description("Get a sports team by id.")
		Response(OK, TeamMedia)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(TeamPayload)
		Params(func() {
			Param("id", String, "Team ID")
		})
		Description("Update a sports team by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Team ID")
		})
		Description("Delete a sports team by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all teams")
		Response(OK, "application/json")
	})
})

var TeamPayload = Type("TeamPayload", func() {
	Description("Team Description.")

	Attribute("name", String, "Team name", func() {
		Metadata("struct:tag:datastore", "name,noindex")
		Metadata("struct:tag:json", "name,omitempty")
	})
	Attribute("sportId", String, "Sport ID", func() {
		Metadata("struct:tag:datastore", "sportId,noindex")
		Metadata("struct:tag:json", "sportId,omitempty")
	})
	Attribute("homeTownId", String, "Sport HomeTown ID", func() {
		Metadata("struct:tag:datastore", "homeTownId,noindex")
		Metadata("struct:tag:json", "homeTownId,omitempty")
	})
	Attribute("shortName", String, "Team Nickname", func() {
		Metadata("struct:tag:datastore", "shortName,noindex")
		Metadata("struct:tag:json", "shortName,omitempty")
	})

	Required("name", "sportId", "homeTownId", "shortName")
})

var TeamMedia = MediaType("application/nmgapi.teamentity", func() {
	Description("Team sport response")
	TypeName("TeamMedia")
	ContentType("application/json")
	Reference(TeamPayload)

	Attributes(func() {
		Attribute("id", String, "ID")
		Attribute("name")
		Attribute("sportId")
		Attribute("homeTownId")
		Attribute("shortName")

		Required("id", "name", "sportId", "homeTownId", "shortName")
	})

	View("default", func() {
		Attribute("name")
		Attribute("sportId")
		Attribute("homeTownId")
		Attribute("shortName")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new teams.")
		Attribute("id")
	})
})
