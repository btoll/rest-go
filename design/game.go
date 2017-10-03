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
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
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
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all games")
		Response(OK, CollectionOf(GameMedia))
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Game ID")
		})
		Description("Get a game by id.")
		Response(OK, GameMedia)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
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
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)
	})
})

var GamePayload = Type("GamePayload", func() {
	Description("Game Description.")

	Attribute("favTeamId", String, "Favorite team id", func() {
		Metadata("struct:tag:datastore", "favTeamId,noindex")
		Metadata("struct:tag:json", "favTeamId,omitempty")
	})
	Attribute("underTeamId", String, "", func() {
		Metadata("struct:tag:datastore", "underTeamId,noindex")
		Metadata("struct:tag:json", "underTeamId,omitempty")
	})
	Attribute("winnerTeamId", String, "Empty until game completed", func() {
		Metadata("struct:tag:datastore", "winnerTeamId,noindex")
		Metadata("struct:tag:json", "winnerTeamId,omitempty")
	})
	Attribute("sportId", String, "Sport ID", func() {
		Metadata("struct:tag:datastore", "sportId,noindex")
		Metadata("struct:tag:json", "sportId,omitempty")
	})
	Attribute("eventId", String, "Event ID", func() {
		Metadata("struct:tag:datastore", "eventId,noindex")
		Metadata("struct:tag:json", "eventId,omitempty")
	})
	Attribute("playDtTm", DateTime, "", func() {
		Metadata("struct:tag:datastore", "playDtTm,noindex")
		Metadata("struct:tag:json", "playDtTm,omitempty")
	})
	Attribute("externalId", String, "Any public id for this game, not unique", func() {
		Metadata("struct:tag:datastore", "externalId,noindex")
		Metadata("struct:tag:json", "externalId,omitempty")
	})
	Attribute("title", String, "Public free form name", func() {
		Metadata("struct:tag:datastore", "title,noindex")
		Metadata("struct:tag:json", "title,omitempty")
	})
	Attribute("locationId", String, "True GPS location", func() {
		Metadata("struct:tag:datastore", "locationId,noindex")
		Metadata("struct:tag:json", "locationId,omitempty")
	})
	Attribute("location", String, "Name of location", func() {
		Metadata("struct:tag:datastore", "location,noindex")
		Metadata("struct:tag:json", "location,omitempty")
	})
	Attribute("oddsForFav", Number, "0", func() {
		Metadata("struct:tag:datastore", "oddsForFav,noindex")
		Metadata("struct:tag:json", "oddsForFav,omitempty")
	})
	Attribute("gameStatus", String, "TeamGameStatus.preGame || tradeable || gameOn || ended", func() {
		Metadata("struct:tag:datastore", "gameStatus,noindex")
		Metadata("struct:tag:json", "gameStatus,omitempty")
	})
	Attribute("finishedAtDtTm", DateTime, "", func() {
		Metadata("struct:tag:datastore", "finishedAtDtTm,noindex")
		Metadata("struct:tag:json", "finishedAtDtTm,omitempty")
	})
	Attribute("loserProgressStatus", String, "TeamGameStatus.eliminated", func() {
		Metadata("struct:tag:datastore", "loserProgressStatus,noindex")
		Metadata("struct:tag:json", "loserProgressStatus,omitempty")
	})

	Required("favTeamId", "underTeamId", "winnerTeamId", "sportId", "playDtTm", "externalId", "title", "locationId", "location", "oddsForFav", "gameStatus", "finishedAtDtTm", "loserProgressStatus")
})

var GameMedia = MediaType("application/nmgapi.gameentity", func() {
	Description("Game response")
	TypeName("GameMedia")
	ContentType("application/json")
	Reference(GamePayload)

	Attributes(func() {
		Attribute("id", String, "ID")
		Attribute("favTeamId")
		Attribute("underTeamId")
		Attribute("winnerTeamId")
		Attribute("sportId")
		Attribute("playDtTm")
		Attribute("externalId")
		Attribute("title")
		Attribute("locationId")
		Attribute("location")
		Attribute("oddsForFav")
		Attribute("gameStatus")
		Attribute("finishedAtDtTm")
		Attribute("loserProgressStatus")

		Required("id", "favTeamId", "underTeamId", "winnerTeamId", "sportId", "playDtTm", "externalId", "title", "locationId", "location", "oddsForFav", "gameStatus", "finishedAtDtTm", "loserProgressStatus")
	})

	View("default", func() {
		Attribute("id")
		Attribute("favTeamId")
		Attribute("underTeamId")
		Attribute("winnerTeamId")
		Attribute("sportId")
		Attribute("playDtTm")
		Attribute("externalId")
		Attribute("title")
		Attribute("locationId")
		Attribute("location")
		Attribute("oddsForFav")
		Attribute("gameStatus")
		Attribute("finishedAtDtTm")
		Attribute("loserProgressStatus")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new teams.")
		Attribute("id")
	})
})
