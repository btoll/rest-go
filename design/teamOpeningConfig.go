package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("TeamOpeningConfig", func() {
	BasePath("/teamOpeningConfig")
	// Seems that goa doesn't like setting DefaultMedia here at the top-level when the MediaType has multiple Views.
	//	DefaultMedia(TeamOpeningConfigMedia)
	Description("Describes a team opening config.")

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new team opening config.")
		Payload(TeamOpeningConfigPayload)
		Response(OK, func() {
			Status(200)
			Media(TeamOpeningConfigMedia, "tiny")
		})
	})

	Action("show", func() {
		//		Routing(GET("/:teamGameEventKey"))
		Routing(GET("/:id"))
		Params(func() {
			//			Param("teamGameEventKey", String, "Team Game Event Key")
			Param("id", String, "Team Game Event Key")
		})
		Description("Get a sports team by event key.")
		Response(OK, TeamOpeningConfigMedia)
	})

	Action("update", func() {
		//		Routing(PUT("/:teamGameEventKey"))
		Routing(PUT("/:id"))
		Payload(TeamOpeningConfigPayload)
		Params(func() {
			//			Param("teamGameEventKey", String, "Team Game Event Key")
			Param("id", String, "Team Game Event Key")
		})
		Description("Update a sports team by event key.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("delete", func() {
		//		Routing(DELETE("/:teamGameEventKey"))
		Routing(DELETE("/:id"))
		Params(func() {
			//			Param("teamGameEventKey", String, "Team Game Event Key")
			Param("id", String, "Team Game Event Key")
		})
		Description("Delete a sports team by event key.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all teams openings")
		Response(OK, "application/json")
	})
})

var TeamOpeningConfigPayload = Type("TeamOpeningConfigPayload", func() {
	Description("Team Description.")

	/*
		Attribute("teamGameEventKey", String, "", func() {
			Metadata("struct:tag:datastore", "__key__")
		})
	*/
	/*
		Attribute("teamGameEventKey", String, "Reference ID for new team", func() {
			Metadata("struct:tag:datastore", "teamGameEventKey")
			Metadata("struct:tag:json", "teamGameEventKey")
		})
	*/
	Attribute("openingPrice", Number, "", func() {
		Metadata("struct:tag:datastore", "openingPrice,noindex")
		Metadata("struct:tag:json", "openingPrice,omitempty")
	})
	Attribute("openingShares", Integer, "", func() {
		Metadata("struct:tag:datastore", "openingShares,noindex")
		Metadata("struct:tag:json", "openingShares,omitempty")
	})
	Attribute("buyIncrementQuan", Integer, "", func() {
		Metadata("struct:tag:datastore", "buyIncrementQuan,noindex")
		Metadata("struct:tag:json", "buyIncrementQuan,omitempty")
	})
	Attribute("buyIncrementPrice", Number, "", func() {
		Metadata("struct:tag:datastore", "buyIncrementPrice,noindex")
		Metadata("struct:tag:json", "buyIncrementPrice,omitempty")
	})
	Attribute("sellDecrementQuan", Integer, "", func() {
		Metadata("struct:tag:datastore", "sellDecrementQuan,noindex")
		Metadata("struct:tag:json", "sellDecrementQuan,omitempty")
	})
	Attribute("sellDecrementPrice", Number, "", func() {
		Metadata("struct:tag:datastore", "sellDecrementPrice,noindex")
		Metadata("struct:tag:json", "sellDecrementPrice,omitempty")
	})
	Attribute("liquidationFee", Number, "", func() {
		Metadata("struct:tag:datastore", "liquidationFee,noindex")
		Metadata("struct:tag:json", "liquidationFee,omitempty")
	})
	Attribute("startTradeDtTm", DateTime, "", func() {
		Metadata("struct:tag:datastore", "startTradeDtTm,noindex")
		Metadata("struct:tag:json", "startTradeDtTm,omitempty")
	})

	//	Required("teamGameEventKey", "openingPrice", "openingShares", "buyIncrementQuan", "buyIncrementPrice", "sellDecrementQuan", "sellDecrementPrice", "liquidationFee", "startTradeDtTm")
	Required("openingPrice", "openingShares", "buyIncrementQuan", "buyIncrementPrice", "sellDecrementQuan", "sellDecrementPrice", "liquidationFee", "startTradeDtTm")
})

var TeamOpeningConfigMedia = MediaType("application/nmgapi.teamopeningconfigentity", func() {
	Description("Team sport response")
	TypeName("TeamOpeningConfigMedia")
	ContentType("application/json")
	Reference(TeamOpeningConfigPayload)

	Attributes(func() {
		//		Attribute("teamGameEventKey")
		Attribute("id", String, "ID")
		Attribute("openingPrice")
		Attribute("openingShares")
		Attribute("buyIncrementQuan")
		Attribute("buyIncrementPrice")
		Attribute("sellDecrementQuan")
		Attribute("sellDecrementPrice")
		Attribute("liquidationFee")
		Attribute("startTradeDtTm")

		//		Required("teamGameEventKey", "openingPrice", "openingShares", "buyIncrementQuan", "buyIncrementPrice", "sellDecrementQuan", "sellDecrementPrice", "liquidationFee", "startTradeDtTm")
		Required("id", "openingPrice", "openingShares", "buyIncrementQuan", "buyIncrementPrice", "sellDecrementQuan", "sellDecrementPrice", "liquidationFee", "startTradeDtTm")
	})

	View("default", func() {
		//		Attribute("teamGameEventKey")
		Attribute("openingPrice")
		Attribute("openingShares")
		Attribute("buyIncrementQuan")
		Attribute("buyIncrementPrice")
		Attribute("sellDecrementQuan")
		Attribute("sellDecrementPrice")
		Attribute("liquidationFee")
		Attribute("startTradeDtTm")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new teams.")
		//		Attribute("teamGameEventKey")
		Attribute("id")
	})
})
