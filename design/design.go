package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("nmgapi", func() {
	Title("New Millennium Games API")
	Description("api for mobile & web clients")
	Host("localhost:8080") // millenniumgames-fa7da.appspot.com
	Scheme("http")
	BasePath("/nmg")
	TermsOfService("nmg tos")
	License(func() { // API Licensing information
		Name("Private (no license offered)")
		URL("http://google.com")
	})
	Docs(func() {
		Description("doc description")
		URL("http://google.com")
	})
})
