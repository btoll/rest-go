//go:generate goagen bootstrap -d github.com/btoll/rest-go/design

package main

import (
	"net/http"

	"github.com/btoll/rest-go/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"google.golang.org/appengine"
)

func main() {
	// Create service
	service := goa.New("nmgapi")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Event" controller
	c := NewEventController(service)
	app.MountEventController(service, c)
	// Mount "Game" controller
	c2 := NewGameController(service)
	app.MountGameController(service, c2)
	// Mount "Sport" controller
	c3 := NewSportController(service)
	app.MountSportController(service, c3)
	// Mount "Team" controller
	c4 := NewTeamController(service)
	app.MountTeamController(service, c4)
	// Mount "TeamOpeningConfig" controller
	c5 := NewTeamOpeningConfigController(service)
	app.MountTeamOpeningConfigController(service, c5)

	//	if t := appengine.IsDevAppServer(); t {
	//		fmt.Printf("appengine.IsDevAppServer() %t\n", t)
	http.Handle("/", service.Mux)
	appengine.Main()
	//	} else {
	//		fmt.Printf("appengine.IsDevAppServer() %t\n", t)
	//
	//		if err := service.ListenAndServe(":8080"); err != nil {
	//			service.LogError("startup", "err", err)
	//		}
	//	}
}
