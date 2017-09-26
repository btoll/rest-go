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
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
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
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all teams")
		Response(OK, CollectionOf(TeamMedia))
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Team ID")
		})
		Description("Get a sports team by id.")
		Response(OK, TeamMedia)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
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
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})
})

var TeamPayload = Type("TeamPayload", func() {
	Description("Team Description.")

	/*
		Attribute("id", String, "Reference ID for new team", func() {
			Metadata("struct:tag:datastore", "__key__")
		})
	*/
	Attribute("id", String, "Reference ID for new team", func() {
		Metadata("struct:tag:datastore", "id")
		Metadata("struct:tag:json", "id")
	})
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
	Attribute("currentWinRecord", String, "Team Win-Loss Record", func() {
		Metadata("struct:tag:datastore", "currentWinRecord,noindex")
		Metadata("struct:tag:json", "currentWinRecord,omitempty")
	}) //"WON 8-LOST 3"
	Attribute("iconName", String, "Team Icon", func() {
		Metadata("struct:tag:datastore", "iconName,noindex")
		Metadata("struct:tag:json", "iconName,omitempty")
	})
	Attribute("fullLogo", String, "Team Logo", func() {
		Metadata("struct:tag:datastore", "fullLogo,noindex")
		Metadata("struct:tag:json", "fullLogo,omitempty")
	})

	Required("id", "name", "sportId", "homeTownId", "shortName", "currentWinRecord", "iconName", "fullLogo")
})

var TeamMedia = MediaType("application/goa.teamentity", func() {
	Description("Team sport response")
	TypeName("TeamMedia")
	ContentType("application/json")
	Reference(TeamPayload)

	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("sportId")
		Attribute("homeTownId")
		Attribute("shortName")
		Attribute("currentWinRecord")
		Attribute("iconName")
		Attribute("fullLogo")

		Required("id", "name", "sportId", "homeTownId", "shortName", "currentWinRecord", "iconName", "fullLogo")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("sportId")
		Attribute("homeTownId")
		Attribute("shortName")
		Attribute("currentWinRecord")
		Attribute("iconName")
		Attribute("fullLogo")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new teams.")
		Attribute("id")
	})
})
