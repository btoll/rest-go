package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Enum", func() {
	BasePath("/enum")
	Description("Get a map of all enums")

	Action("list", func() {
		Routing(GET("/"))
		Description("Get a map of all enums")
		Response(OK, "application/json")
	})

})
