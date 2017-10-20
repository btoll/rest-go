// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// EventHref returns the resource href.
func EventHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nmg/event/%v", paramid)
}

// GameHref returns the resource href.
func GameHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nmg/game/%v", paramid)
}

// SportHref returns the resource href.
func SportHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nmg/sport/%v", paramid)
}

// TeamHref returns the resource href.
func TeamHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nmg/team/%v", paramid)
}

// TeamOpeningConfigHref returns the resource href.
func TeamOpeningConfigHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/nmg/teamOpeningConfig/%v", paramid)
}
