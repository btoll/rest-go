// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/dgaedcke/nmg_admin_service/design
// --out=$(GOPATH)/src/github.com/dgaedcke/nmg_admin_service
// --regen=true
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// Event sport response (default view)
//
// Identifier: application/goa.evententity; view=default
type EventMedia struct {
	EndDtTm time.Time `datastore:"endDtTm,noindex" json:"endDtTm,omitempty"`
	// Not guaranteed to be unique
	EventID string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Reference ID for new event
	ID string `datastore:"id" json:"id"`
	// Location.defaultLoc.id
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// e.g., March Madness
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Sport ID
	SportID   string    `datastore:"sportId,noindex" json:"sportId,omitempty"`
	StartDtTm time.Time `datastore:"startDtTm,noindex" json:"startDtTm,omitempty"`
	SubTitle  string    `datastore:"subTitle,noindex" json:"subTitle,omitempty"`
	// EventAdvanceMethod.singleElimination || doubleElim || bestOf
	TeamAdvanceMethod string `datastore:"teamAdvanceMethod,noindex" json:"teamAdvanceMethod,omitempty"`
}

// Validate validates the EventMedia media type instance.
func (mt *EventMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}
	if mt.EventID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eventId"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.SubTitle == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subTitle"))
	}

	if mt.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "locationId"))
	}
	if mt.TeamAdvanceMethod == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "teamAdvanceMethod"))
	}
	return
}

// Event sport response (tiny view)
//
// Identifier: application/goa.evententity; view=tiny
type EventMediaTiny struct {
	// Reference ID for new event
	ID string `datastore:"id" json:"id"`
}

// Validate validates the EventMediaTiny media type instance.
func (mt *EventMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeEventMedia decodes the EventMedia instance encoded in resp body.
func (c *Client) DecodeEventMedia(resp *http.Response) (*EventMedia, error) {
	var decoded EventMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeEventMediaTiny decodes the EventMediaTiny instance encoded in resp body.
func (c *Client) DecodeEventMediaTiny(resp *http.Response) (*EventMediaTiny, error) {
	var decoded EventMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// EventMediaCollection is the media type for an array of EventMedia (default view)
//
// Identifier: application/goa.evententity; type=collection; view=default
type EventMediaCollection []*EventMedia

// Validate validates the EventMediaCollection media type instance.
func (mt EventMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// EventMediaCollection is the media type for an array of EventMedia (tiny view)
//
// Identifier: application/goa.evententity; type=collection; view=tiny
type EventMediaTinyCollection []*EventMediaTiny

// Validate validates the EventMediaTinyCollection media type instance.
func (mt EventMediaTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeEventMediaCollection decodes the EventMediaCollection instance encoded in resp body.
func (c *Client) DecodeEventMediaCollection(resp *http.Response) (EventMediaCollection, error) {
	var decoded EventMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeEventMediaTinyCollection decodes the EventMediaTinyCollection instance encoded in resp body.
func (c *Client) DecodeEventMediaTinyCollection(resp *http.Response) (EventMediaTinyCollection, error) {
	var decoded EventMediaTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Game response (default view)
//
// Identifier: application/goa.gameentity; view=default
type GameMedia struct {
	// Any public id for this game, not unique
	ExternalID string `datastore:"externalId,noindex" json:"externalId,omitempty"`
	// Favorite team id
	FavTeamID      string    `datastore:"favTeamId,noindex" json:"favTeamId,omitempty"`
	FinishedAtDtTm time.Time `datastore:"finishedAtDtTm,noindex" json:"finishedAtDtTm,omitempty"`
	// TeamGameStatus.preGame || tradeable || gameOn || ended
	GameStatus string `datastore:"gameStatus,noindex" json:"gameStatus,omitempty"`
	// Reference ID for new team
	ID string `datastore:"id" json:"id"`
	// Name of location
	Location string `datastore:"location,noindex" json:"location,omitempty"`
	// True GPS location
	LocationID string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// TeamGameStatus.eliminated
	LoserProgressStatus string `datastore:"loserProgressStatus,noindex" json:"loserProgressStatus,omitempty"`
	// 0
	OddsForFav float64   `datastore:"oddsForFav,noindex" json:"oddsForFav,omitempty"`
	PlayDtTm   time.Time `datastore:"playDtTm,noindex" json:"playDtTm,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
	// Public free form name
	Title       string `datastore:"title,noindex" json:"title,omitempty"`
	UnderTeamID string `datastore:"underTeamId,noindex" json:"underTeamId,omitempty"`
	// Empty until game completed
	WinnerTeamID string `datastore:"winnerTeamId,noindex" json:"winnerTeamId,omitempty"`
}

// Validate validates the GameMedia media type instance.
func (mt *GameMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.FavTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "favTeamId"))
	}
	if mt.UnderTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "underTeamId"))
	}
	if mt.WinnerTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "winnerTeamId"))
	}
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}

	if mt.ExternalID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "externalId"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "locationId"))
	}
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if mt.GameStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "gameStatus"))
	}

	if mt.LoserProgressStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "loserProgressStatus"))
	}
	return
}

// Game response (tiny view)
//
// Identifier: application/goa.gameentity; view=tiny
type GameMediaTiny struct {
	// Reference ID for new team
	ID string `datastore:"id" json:"id"`
}

// Validate validates the GameMediaTiny media type instance.
func (mt *GameMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeGameMedia decodes the GameMedia instance encoded in resp body.
func (c *Client) DecodeGameMedia(resp *http.Response) (*GameMedia, error) {
	var decoded GameMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeGameMediaTiny decodes the GameMediaTiny instance encoded in resp body.
func (c *Client) DecodeGameMediaTiny(resp *http.Response) (*GameMediaTiny, error) {
	var decoded GameMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GameMediaCollection is the media type for an array of GameMedia (default view)
//
// Identifier: application/goa.gameentity; type=collection; view=default
type GameMediaCollection []*GameMedia

// Validate validates the GameMediaCollection media type instance.
func (mt GameMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GameMediaCollection is the media type for an array of GameMedia (tiny view)
//
// Identifier: application/goa.gameentity; type=collection; view=tiny
type GameMediaTinyCollection []*GameMediaTiny

// Validate validates the GameMediaTinyCollection media type instance.
func (mt GameMediaTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeGameMediaCollection decodes the GameMediaCollection instance encoded in resp body.
func (c *Client) DecodeGameMediaCollection(resp *http.Response) (GameMediaCollection, error) {
	var decoded GameMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeGameMediaTinyCollection decodes the GameMediaTinyCollection instance encoded in resp body.
func (c *Client) DecodeGameMediaTinyCollection(resp *http.Response) (GameMediaTinyCollection, error) {
	var decoded GameMediaTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Sport response (default view)
//
// Identifier: application/goa.sportentity; view=default
type SportMedia struct {
	// Is in season?
	Active bool `datastore:"active,noindex" json:"active,omitempty"`
	// sport_icon
	BackgroundImageName string `datastore:"backgroundImageName,noindex" json:"backgroundImageName,omitempty"`
	// Tournament
	EventTerm string `datastore:"eventTerm,noindex" json:"eventTerm,omitempty"`
	// Game
	GameTerm string `datastore:"gameTerm,noindex" json:"gameTerm,omitempty"`
	// sport_icon
	IconName string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Reference ID for new sport
	ID               string  `datastore:"id" json:"id"`
	MaxPreSplitPrice float64 `datastore:"maxPreSplitPrice,noindex" json:"maxPreSplitPrice,omitempty"`
	// Sport name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
}

// Validate validates the SportMedia media type instance.
func (mt *SportMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.GameTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "gameTerm"))
	}
	if mt.EventTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "eventTerm"))
	}
	if mt.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "iconName"))
	}
	if mt.BackgroundImageName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "backgroundImageName"))
	}
	return
}

// Sport response (tiny view)
//
// Identifier: application/goa.sportentity; view=tiny
type SportMediaTiny struct {
	// Reference ID for new sport
	ID string `datastore:"id" json:"id"`
}

// Validate validates the SportMediaTiny media type instance.
func (mt *SportMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeSportMedia decodes the SportMedia instance encoded in resp body.
func (c *Client) DecodeSportMedia(resp *http.Response) (*SportMedia, error) {
	var decoded SportMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeSportMediaTiny decodes the SportMediaTiny instance encoded in resp body.
func (c *Client) DecodeSportMediaTiny(resp *http.Response) (*SportMediaTiny, error) {
	var decoded SportMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// SportMediaCollection is the media type for an array of SportMedia (default view)
//
// Identifier: application/goa.sportentity; type=collection; view=default
type SportMediaCollection []*SportMedia

// Validate validates the SportMediaCollection media type instance.
func (mt SportMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// SportMediaCollection is the media type for an array of SportMedia (tiny view)
//
// Identifier: application/goa.sportentity; type=collection; view=tiny
type SportMediaTinyCollection []*SportMediaTiny

// Validate validates the SportMediaTinyCollection media type instance.
func (mt SportMediaTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeSportMediaCollection decodes the SportMediaCollection instance encoded in resp body.
func (c *Client) DecodeSportMediaCollection(resp *http.Response) (SportMediaCollection, error) {
	var decoded SportMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeSportMediaTinyCollection decodes the SportMediaTinyCollection instance encoded in resp body.
func (c *Client) DecodeSportMediaTinyCollection(resp *http.Response) (SportMediaTinyCollection, error) {
	var decoded SportMediaTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Team sport response (default view)
//
// Identifier: application/goa.teamentity; view=default
type TeamMedia struct {
	// Team Win-Loss Record
	CurrentWinRecord string `datastore:"currentWinRecord,noindex" json:"currentWinRecord,omitempty"`
	// Team Logo
	FullLogo string `datastore:"fullLogo,noindex" json:"fullLogo,omitempty"`
	// Sport HomeTown ID
	HomeTownID string `datastore:"homeTownId,noindex" json:"homeTownId,omitempty"`
	// Team Icon
	IconName string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Reference ID for new team
	ID string `datastore:"id" json:"id"`
	// Team name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Team Nickname
	ShortName string `datastore:"shortName,noindex" json:"shortName,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
}

// Validate validates the TeamMedia media type instance.
func (mt *TeamMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sportId"))
	}
	if mt.HomeTownID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "homeTownId"))
	}
	if mt.ShortName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "shortName"))
	}
	if mt.CurrentWinRecord == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "currentWinRecord"))
	}
	if mt.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "iconName"))
	}
	if mt.FullLogo == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fullLogo"))
	}
	return
}

// Team sport response (tiny view)
//
// Identifier: application/goa.teamentity; view=tiny
type TeamMediaTiny struct {
	// Reference ID for new team
	ID string `datastore:"id" json:"id"`
}

// Validate validates the TeamMediaTiny media type instance.
func (mt *TeamMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeTeamMedia decodes the TeamMedia instance encoded in resp body.
func (c *Client) DecodeTeamMedia(resp *http.Response) (*TeamMedia, error) {
	var decoded TeamMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeTeamMediaTiny decodes the TeamMediaTiny instance encoded in resp body.
func (c *Client) DecodeTeamMediaTiny(resp *http.Response) (*TeamMediaTiny, error) {
	var decoded TeamMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamMediaCollection is the media type for an array of TeamMedia (default view)
//
// Identifier: application/goa.teamentity; type=collection; view=default
type TeamMediaCollection []*TeamMedia

// Validate validates the TeamMediaCollection media type instance.
func (mt TeamMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// TeamMediaCollection is the media type for an array of TeamMedia (tiny view)
//
// Identifier: application/goa.teamentity; type=collection; view=tiny
type TeamMediaTinyCollection []*TeamMediaTiny

// Validate validates the TeamMediaTinyCollection media type instance.
func (mt TeamMediaTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamMediaCollection decodes the TeamMediaCollection instance encoded in resp body.
func (c *Client) DecodeTeamMediaCollection(resp *http.Response) (TeamMediaCollection, error) {
	var decoded TeamMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeTeamMediaTinyCollection decodes the TeamMediaTinyCollection instance encoded in resp body.
func (c *Client) DecodeTeamMediaTinyCollection(resp *http.Response) (TeamMediaTinyCollection, error) {
	var decoded TeamMediaTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Team sport response (default view)
//
// Identifier: application/goa.teamopeningconfigentity; view=default
type TeamOpeningConfigMedia struct {
	BuyIncrementPrice float64 `datastore:"buyIncrementPrice,noindex" json:"buyIncrementPrice,omitempty"`
	BuyIncrementQuan  int     `datastore:"buyIncrementQuan,noindex" json:"buyIncrementQuan,omitempty"`
	// Reference ID for new team
	ID                 string    `datastore:"id" json:"id"`
	LiquidationFee     float64   `datastore:"liquidationFee,noindex" json:"liquidationFee,omitempty"`
	OpeningPrice       float64   `datastore:"openingPrice,noindex" json:"openingPrice,omitempty"`
	OpeningShares      int       `datastore:"openingShares,noindex" json:"openingShares,omitempty"`
	SellDecrementPrice float64   `datastore:"sellDecrementPrice,noindex" json:"sellDecrementPrice,omitempty"`
	SellDecrementQuan  int       `datastore:"sellDecrementQuan,noindex" json:"sellDecrementQuan,omitempty"`
	StartTradeDtTm     time.Time `datastore:"startTradeDtTm,noindex" json:"startTradeDtTm,omitempty"`
}

// Validate validates the TeamOpeningConfigMedia media type instance.
func (mt *TeamOpeningConfigMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	return
}

// Team sport response (tiny view)
//
// Identifier: application/goa.teamopeningconfigentity; view=tiny
type TeamOpeningConfigMediaTiny struct {
	// Reference ID for new team
	ID string `datastore:"id" json:"id"`
}

// Validate validates the TeamOpeningConfigMediaTiny media type instance.
func (mt *TeamOpeningConfigMediaTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeTeamOpeningConfigMedia decodes the TeamOpeningConfigMedia instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMedia(resp *http.Response) (*TeamOpeningConfigMedia, error) {
	var decoded TeamOpeningConfigMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeTeamOpeningConfigMediaTiny decodes the TeamOpeningConfigMediaTiny instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMediaTiny(resp *http.Response) (*TeamOpeningConfigMediaTiny, error) {
	var decoded TeamOpeningConfigMediaTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamOpeningConfigMediaCollection is the media type for an array of TeamOpeningConfigMedia (default view)
//
// Identifier: application/goa.teamopeningconfigentity; type=collection; view=default
type TeamOpeningConfigMediaCollection []*TeamOpeningConfigMedia

// Validate validates the TeamOpeningConfigMediaCollection media type instance.
func (mt TeamOpeningConfigMediaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// TeamOpeningConfigMediaCollection is the media type for an array of TeamOpeningConfigMedia (tiny view)
//
// Identifier: application/goa.teamopeningconfigentity; type=collection; view=tiny
type TeamOpeningConfigMediaTinyCollection []*TeamOpeningConfigMediaTiny

// Validate validates the TeamOpeningConfigMediaTinyCollection media type instance.
func (mt TeamOpeningConfigMediaTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamOpeningConfigMediaCollection decodes the TeamOpeningConfigMediaCollection instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMediaCollection(resp *http.Response) (TeamOpeningConfigMediaCollection, error) {
	var decoded TeamOpeningConfigMediaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeTeamOpeningConfigMediaTinyCollection decodes the TeamOpeningConfigMediaTinyCollection instance encoded in resp body.
func (c *Client) DecodeTeamOpeningConfigMediaTinyCollection(resp *http.Response) (TeamOpeningConfigMediaTinyCollection, error) {
	var decoded TeamOpeningConfigMediaTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}