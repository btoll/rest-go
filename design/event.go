package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Event", func() {
	BasePath("/event")
	// Seems that goa doesn't like setting DefaultMedia here at the top-level when the MediaType has multiple Views.
	//	DefaultMedia(EventMedia)
	Description("Describes a sport event.")

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new sports event.")
		Payload(EventPayload)
		Response(OK, func() {
			Status(200)
			Media(EventMedia, "tiny")
		})
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Event ID")
		})
		Description("Get a sports event by id.")
		Response(OK, EventMedia)
	})

	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(EventPayload)
		Params(func() {
			Param("id", String, "Event ID")
		})
		Description("Update a sports event by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Event ID")
		})
		Description("Delete a sports event by id.")
		Response(OK, func() {
			Status(200)
		})
		Response(NoContent)
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("Get all events")
		Response(OK, "application/json")
	})
})

var EventPayload = Type("EventPayload", func() {
	Description("Event Description.")

	Attribute("sportId", String, "Sport ID", func() {
		Metadata("struct:tag:datastore", "sportId,noindex")
		Metadata("struct:tag:json", "sportId,omitempty")
	})
	Attribute("eventId", String, "Not guaranteed to be unique", func() {
		Metadata("struct:tag:datastore", "eventId,noindex")
		Metadata("struct:tag:json", "eventId,omitempty")
	})
	Attribute("name", String, "e.g., March Madness", func() {
		Metadata("struct:tag:datastore", "name,noindex")
		Metadata("struct:tag:json", "name,omitempty")
	})
	Attribute("subTitle", String, "", func() {
		Metadata("struct:tag:datastore", "subTitle,noindex")
		Metadata("struct:tag:json", "subTitle,omitempty")
	})
	Attribute("startDtTm", DateTime, "", func() {
		Metadata("struct:tag:datastore", "startDtTm,noindex")
		Metadata("struct:tag:json", "startDtTm,omitempty")
	})
	Attribute("endDtTm", DateTime, "", func() {
		Metadata("struct:tag:datastore", "endDtTm,noindex")
		Metadata("struct:tag:json", "endDtTm,omitempty")
	})
	Attribute("locationId", String, "Location.defaultLoc.id", func() {
		Metadata("struct:tag:datastore", "locationId,noindex")
		Metadata("struct:tag:json", "locationId,omitempty")
	})
	Attribute("teamAdvanceMethod", String, "EventAdvanceMethod.singleElimination || doubleElim || bestOf", func() {
		Metadata("struct:tag:datastore", "teamAdvanceMethod,noindex")
		Metadata("struct:tag:json", "teamAdvanceMethod,omitempty")
	})

	Required("sportId", "eventId", "name", "subTitle", "startDtTm", "endDtTm", "locationId", "teamAdvanceMethod")
})

var EventMedia = MediaType("application/nmgapi.evententity", func() {
	Description("Event sport response")
	TypeName("EventMedia")
	ContentType("application/json")
	Reference(EventPayload)

	Attributes(func() {
		Attribute("id", String, "ID")
		Attribute("sportId")
		Attribute("eventId")
		Attribute("name")
		Attribute("subTitle")
		Attribute("startDtTm")
		Attribute("endDtTm")
		Attribute("locationId")
		Attribute("teamAdvanceMethod")

		Required("id", "sportId", "eventId", "name", "subTitle", "startDtTm", "endDtTm", "locationId", "teamAdvanceMethod")
	})

	View("default", func() {
		Attribute("sportId")
		Attribute("eventId")
		Attribute("name")
		Attribute("subTitle")
		Attribute("startDtTm")
		Attribute("endDtTm")
		Attribute("locationId")
		Attribute("teamAdvanceMethod")
	})

	View("tiny", func() {
		Description("`tiny` is the view used to create new events.")
		Attribute("id")
	})
})
