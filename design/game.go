package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Game", func() {
	BasePath("/game")
	// Seems that goa doesn't like setting DefaultMedia here at the top-level when the MediaType has multiple Views.
	//	DefaultMedia(TeamMedia)
	Description("Describes a game.")

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new game.")
		Payload(GamePayload)
		Response(OK, func() {
			Status(200)
			Media(GameMedia, "tiny")
		})
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Game ID")
		})
		Description("Get a game by id.")
		Response(OK, GameMedia)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(GamePayload)
		Params(func() {
			Param("id", String, "Game ID")
		})
		Description("Update a game by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Game ID")
		})
		Description("Delete a game by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all games")
		Response(OK, CollectionOf(GameMedia))
	})
})

var GamePayload = Type("GamePayload", func() {
	Description("Game Description.")

	Attribute("id", String, "ID", func() {
		Metadata("struct:tag:datastore", "id,noindex")
		Metadata("struct:tag:json", "id,omitempty")
	})
	Attribute("sportId", String, "Sport ID", func() {
		Metadata("struct:tag:datastore", "sportId,noindex")
		Metadata("struct:tag:json", "sportId,omitempty")
	})
	Attribute("eventId", String, "Event ID", func() {
		Metadata("struct:tag:datastore", "eventId,noindex")
		Metadata("struct:tag:json", "eventId,omitempty")
	})
	Attribute("gamePlayStatus", String, "TeamGamePlayStatus.preGame || tradeable || gameOn || ended", func() {
		Metadata("struct:tag:datastore", "gamePlayStatus,noindex")
		Metadata("struct:tag:json", "gamePlayStatus,omitempty")
	})
	Attribute("loserAdvanceState", String, "TeamGameStatus.eliminated", func() {
		Metadata("struct:tag:datastore", "loserAdvanceState,noindex")
		Metadata("struct:tag:json", "loserAdvanceState,omitempty")
	})

	Required("id", "sportId", "eventId", "gamePlayStatus", "loserAdvanceState")
})

var GameMedia = MediaType("application/nmgapi.gameentity", func() {
	Description("Game response")
	TypeName("GameMedia")
	ContentType("application/json")
	Reference(GamePayload)

	Attributes(func() {
		Attribute("id")
		Attribute("sportId")
		Attribute("eventId")
		Attribute("gamePlayStatus")
		Attribute("loserAdvanceState")

		Required("id", "sportId", "eventId", "gamePlayStatus", "loserAdvanceState")
	})

	View("default", func() {
		Attribute("id")
		Attribute("sportId")
		Attribute("eventId")
		Attribute("gamePlayStatus")
		Attribute("loserAdvanceState")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new teams.")
		Attribute("id")
	})
})
