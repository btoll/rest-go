// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Application User Types
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --regen=true
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// Event Description.
type eventPayload struct {
	EndDtTm *time.Time `datastore:"endDtTm,noindex" json:"endDtTm,omitempty"`
	// Not guaranteed to be unique
	EventID *string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Location.defaultLoc.id
	LocationID *string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// e.g., March Madness
	Name *string `datastore:"name,noindex" json:"name,omitempty"`
	// Sport ID
	SportID   *string    `datastore:"sportId,noindex" json:"sportId,omitempty"`
	StartDtTm *time.Time `datastore:"startDtTm,noindex" json:"startDtTm,omitempty"`
	SubTitle  *string    `datastore:"subTitle,noindex" json:"subTitle,omitempty"`
	// EventAdvanceMethod.singleElimination || doubleElim || bestOf
	TeamAdvanceMethod *string `datastore:"teamAdvanceMethod,noindex" json:"teamAdvanceMethod,omitempty"`
}

// Validate validates the eventPayload type instance.
func (ut *eventPayload) Validate() (err error) {
	if ut.SportID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sportId"))
	}
	if ut.EventID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "eventId"))
	}
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.SubTitle == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "subTitle"))
	}
	if ut.StartDtTm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "startDtTm"))
	}
	if ut.EndDtTm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "endDtTm"))
	}
	if ut.LocationID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "locationId"))
	}
	if ut.TeamAdvanceMethod == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "teamAdvanceMethod"))
	}
	return
}

// Publicize creates EventPayload from eventPayload
func (ut *eventPayload) Publicize() *EventPayload {
	var pub EventPayload
	if ut.EndDtTm != nil {
		pub.EndDtTm = *ut.EndDtTm
	}
	if ut.EventID != nil {
		pub.EventID = *ut.EventID
	}
	if ut.LocationID != nil {
		pub.LocationID = *ut.LocationID
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.SportID != nil {
		pub.SportID = *ut.SportID
	}
	if ut.StartDtTm != nil {
		pub.StartDtTm = *ut.StartDtTm
	}
	if ut.SubTitle != nil {
		pub.SubTitle = *ut.SubTitle
	}
	if ut.TeamAdvanceMethod != nil {
		pub.TeamAdvanceMethod = *ut.TeamAdvanceMethod
	}
	return &pub
}

// Event Description.
type EventPayload struct {
	EndDtTm time.Time `datastore:"endDtTm,noindex" json:"endDtTm,omitempty"`
	// Not guaranteed to be unique
	EventID string `datastore:"eventId,noindex" json:"eventId,omitempty"`
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

// Validate validates the EventPayload type instance.
func (ut *EventPayload) Validate() (err error) {
	if ut.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sportId"))
	}
	if ut.EventID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "eventId"))
	}
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.SubTitle == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "subTitle"))
	}

	if ut.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "locationId"))
	}
	if ut.TeamAdvanceMethod == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "teamAdvanceMethod"))
	}
	return
}

// Game Description.
type gamePayload struct {
	// Event ID
	EventID *string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Any public id for this game, not unique
	ExternalID *string `datastore:"externalId,noindex" json:"externalId,omitempty"`
	// Favorite team id
	FavTeamID      *string    `datastore:"favTeamId,noindex" json:"favTeamId,omitempty"`
	FinishedAtDtTm *time.Time `datastore:"finishedAtDtTm,noindex" json:"finishedAtDtTm,omitempty"`
	// TeamGameStatus.preGame || tradeable || gameOn || ended
	GameStatus *string `datastore:"gameStatus,noindex" json:"gameStatus,omitempty"`
	// Name of location
	Location *string `datastore:"location,noindex" json:"location,omitempty"`
	// True GPS location
	LocationID *string `datastore:"locationId,noindex" json:"locationId,omitempty"`
	// TeamGameStatus.eliminated
	LoserProgressStatus *string `datastore:"loserProgressStatus,noindex" json:"loserProgressStatus,omitempty"`
	// 0
	OddsForFav *float64   `datastore:"oddsForFav,noindex" json:"oddsForFav,omitempty"`
	PlayDtTm   *time.Time `datastore:"playDtTm,noindex" json:"playDtTm,omitempty"`
	// Sport ID
	SportID *string `datastore:"sportId,noindex" json:"sportId,omitempty"`
	// Public free form name
	Title       *string `datastore:"title,noindex" json:"title,omitempty"`
	UnderTeamID *string `datastore:"underTeamId,noindex" json:"underTeamId,omitempty"`
	// Empty until game completed
	WinnerTeamID *string `datastore:"winnerTeamId,noindex" json:"winnerTeamId,omitempty"`
}

// Validate validates the gamePayload type instance.
func (ut *gamePayload) Validate() (err error) {
	if ut.FavTeamID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "favTeamId"))
	}
	if ut.UnderTeamID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "underTeamId"))
	}
	if ut.WinnerTeamID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "winnerTeamId"))
	}
	if ut.SportID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sportId"))
	}
	if ut.PlayDtTm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "playDtTm"))
	}
	if ut.ExternalID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "externalId"))
	}
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.LocationID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "locationId"))
	}
	if ut.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "location"))
	}
	if ut.OddsForFav == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "oddsForFav"))
	}
	if ut.GameStatus == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "gameStatus"))
	}
	if ut.FinishedAtDtTm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "finishedAtDtTm"))
	}
	if ut.LoserProgressStatus == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "loserProgressStatus"))
	}
	return
}

// Publicize creates GamePayload from gamePayload
func (ut *gamePayload) Publicize() *GamePayload {
	var pub GamePayload
	if ut.EventID != nil {
		pub.EventID = ut.EventID
	}
	if ut.ExternalID != nil {
		pub.ExternalID = *ut.ExternalID
	}
	if ut.FavTeamID != nil {
		pub.FavTeamID = *ut.FavTeamID
	}
	if ut.FinishedAtDtTm != nil {
		pub.FinishedAtDtTm = *ut.FinishedAtDtTm
	}
	if ut.GameStatus != nil {
		pub.GameStatus = *ut.GameStatus
	}
	if ut.Location != nil {
		pub.Location = *ut.Location
	}
	if ut.LocationID != nil {
		pub.LocationID = *ut.LocationID
	}
	if ut.LoserProgressStatus != nil {
		pub.LoserProgressStatus = *ut.LoserProgressStatus
	}
	if ut.OddsForFav != nil {
		pub.OddsForFav = *ut.OddsForFav
	}
	if ut.PlayDtTm != nil {
		pub.PlayDtTm = *ut.PlayDtTm
	}
	if ut.SportID != nil {
		pub.SportID = *ut.SportID
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	if ut.UnderTeamID != nil {
		pub.UnderTeamID = *ut.UnderTeamID
	}
	if ut.WinnerTeamID != nil {
		pub.WinnerTeamID = *ut.WinnerTeamID
	}
	return &pub
}

// Game Description.
type GamePayload struct {
	// Event ID
	EventID *string `datastore:"eventId,noindex" json:"eventId,omitempty"`
	// Any public id for this game, not unique
	ExternalID string `datastore:"externalId,noindex" json:"externalId,omitempty"`
	// Favorite team id
	FavTeamID      string    `datastore:"favTeamId,noindex" json:"favTeamId,omitempty"`
	FinishedAtDtTm time.Time `datastore:"finishedAtDtTm,noindex" json:"finishedAtDtTm,omitempty"`
	// TeamGameStatus.preGame || tradeable || gameOn || ended
	GameStatus string `datastore:"gameStatus,noindex" json:"gameStatus,omitempty"`
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

// Validate validates the GamePayload type instance.
func (ut *GamePayload) Validate() (err error) {
	if ut.FavTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "favTeamId"))
	}
	if ut.UnderTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "underTeamId"))
	}
	if ut.WinnerTeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "winnerTeamId"))
	}
	if ut.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sportId"))
	}

	if ut.ExternalID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "externalId"))
	}
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.LocationID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "locationId"))
	}
	if ut.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "location"))
	}

	if ut.GameStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "gameStatus"))
	}

	if ut.LoserProgressStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "loserProgressStatus"))
	}
	return
}

// Sport Description.
type sportPayload struct {
	// Is in season?
	Active *bool `datastore:"active,noindex" json:"active,omitempty"`
	// sport_icon
	BackgroundImageName *string `datastore:"backgroundImageName,noindex" json:"backgroundImageName,omitempty"`
	// Tournament
	EventTerm *string `datastore:"eventTerm,noindex" json:"eventTerm,omitempty"`
	// Game
	GameTerm *string `datastore:"gameTerm,noindex" json:"gameTerm,omitempty"`
	// sport_icon
	IconName         *string  `datastore:"iconName,noindex" json:"iconName,omitempty"`
	MaxPreSplitPrice *float64 `datastore:"maxPreSplitPrice,noindex" json:"maxPreSplitPrice,omitempty"`
	// Sport name
	Name *string `datastore:"name,noindex" json:"name,omitempty"`
}

// Validate validates the sportPayload type instance.
func (ut *sportPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Active == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "active"))
	}
	if ut.MaxPreSplitPrice == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "maxPreSplitPrice"))
	}
	if ut.GameTerm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "gameTerm"))
	}
	if ut.EventTerm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "eventTerm"))
	}
	if ut.IconName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "iconName"))
	}
	if ut.BackgroundImageName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "backgroundImageName"))
	}
	return
}

// Publicize creates SportPayload from sportPayload
func (ut *sportPayload) Publicize() *SportPayload {
	var pub SportPayload
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.BackgroundImageName != nil {
		pub.BackgroundImageName = *ut.BackgroundImageName
	}
	if ut.EventTerm != nil {
		pub.EventTerm = *ut.EventTerm
	}
	if ut.GameTerm != nil {
		pub.GameTerm = *ut.GameTerm
	}
	if ut.IconName != nil {
		pub.IconName = *ut.IconName
	}
	if ut.MaxPreSplitPrice != nil {
		pub.MaxPreSplitPrice = *ut.MaxPreSplitPrice
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// Sport Description.
type SportPayload struct {
	// Is in season?
	Active bool `datastore:"active,noindex" json:"active,omitempty"`
	// sport_icon
	BackgroundImageName string `datastore:"backgroundImageName,noindex" json:"backgroundImageName,omitempty"`
	// Tournament
	EventTerm string `datastore:"eventTerm,noindex" json:"eventTerm,omitempty"`
	// Game
	GameTerm string `datastore:"gameTerm,noindex" json:"gameTerm,omitempty"`
	// sport_icon
	IconName         string  `datastore:"iconName,noindex" json:"iconName,omitempty"`
	MaxPreSplitPrice float64 `datastore:"maxPreSplitPrice,noindex" json:"maxPreSplitPrice,omitempty"`
	// Sport name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
}

// Validate validates the SportPayload type instance.
func (ut *SportPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}

	if ut.GameTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "gameTerm"))
	}
	if ut.EventTerm == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "eventTerm"))
	}
	if ut.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "iconName"))
	}
	if ut.BackgroundImageName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "backgroundImageName"))
	}
	return
}

// Team Description.
type teamOpeningConfigPayload struct {
	BuyIncrementPrice  *float64   `datastore:"buyIncrementPrice,noindex" json:"buyIncrementPrice,omitempty"`
	BuyIncrementQuan   *int       `datastore:"buyIncrementQuan,noindex" json:"buyIncrementQuan,omitempty"`
	LiquidationFee     *float64   `datastore:"liquidationFee,noindex" json:"liquidationFee,omitempty"`
	OpeningPrice       *float64   `datastore:"openingPrice,noindex" json:"openingPrice,omitempty"`
	OpeningShares      *int       `datastore:"openingShares,noindex" json:"openingShares,omitempty"`
	SellDecrementPrice *float64   `datastore:"sellDecrementPrice,noindex" json:"sellDecrementPrice,omitempty"`
	SellDecrementQuan  *int       `datastore:"sellDecrementQuan,noindex" json:"sellDecrementQuan,omitempty"`
	StartTradeDtTm     *time.Time `datastore:"startTradeDtTm,noindex" json:"startTradeDtTm,omitempty"`
}

// Validate validates the teamOpeningConfigPayload type instance.
func (ut *teamOpeningConfigPayload) Validate() (err error) {
	if ut.OpeningPrice == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "openingPrice"))
	}
	if ut.OpeningShares == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "openingShares"))
	}
	if ut.BuyIncrementQuan == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "buyIncrementQuan"))
	}
	if ut.BuyIncrementPrice == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "buyIncrementPrice"))
	}
	if ut.SellDecrementQuan == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sellDecrementQuan"))
	}
	if ut.SellDecrementPrice == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sellDecrementPrice"))
	}
	if ut.LiquidationFee == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "liquidationFee"))
	}
	if ut.StartTradeDtTm == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "startTradeDtTm"))
	}
	return
}

// Publicize creates TeamOpeningConfigPayload from teamOpeningConfigPayload
func (ut *teamOpeningConfigPayload) Publicize() *TeamOpeningConfigPayload {
	var pub TeamOpeningConfigPayload
	if ut.BuyIncrementPrice != nil {
		pub.BuyIncrementPrice = *ut.BuyIncrementPrice
	}
	if ut.BuyIncrementQuan != nil {
		pub.BuyIncrementQuan = *ut.BuyIncrementQuan
	}
	if ut.LiquidationFee != nil {
		pub.LiquidationFee = *ut.LiquidationFee
	}
	if ut.OpeningPrice != nil {
		pub.OpeningPrice = *ut.OpeningPrice
	}
	if ut.OpeningShares != nil {
		pub.OpeningShares = *ut.OpeningShares
	}
	if ut.SellDecrementPrice != nil {
		pub.SellDecrementPrice = *ut.SellDecrementPrice
	}
	if ut.SellDecrementQuan != nil {
		pub.SellDecrementQuan = *ut.SellDecrementQuan
	}
	if ut.StartTradeDtTm != nil {
		pub.StartTradeDtTm = *ut.StartTradeDtTm
	}
	return &pub
}

// Team Description.
type TeamOpeningConfigPayload struct {
	BuyIncrementPrice  float64   `datastore:"buyIncrementPrice,noindex" json:"buyIncrementPrice,omitempty"`
	BuyIncrementQuan   int       `datastore:"buyIncrementQuan,noindex" json:"buyIncrementQuan,omitempty"`
	LiquidationFee     float64   `datastore:"liquidationFee,noindex" json:"liquidationFee,omitempty"`
	OpeningPrice       float64   `datastore:"openingPrice,noindex" json:"openingPrice,omitempty"`
	OpeningShares      int       `datastore:"openingShares,noindex" json:"openingShares,omitempty"`
	SellDecrementPrice float64   `datastore:"sellDecrementPrice,noindex" json:"sellDecrementPrice,omitempty"`
	SellDecrementQuan  int       `datastore:"sellDecrementQuan,noindex" json:"sellDecrementQuan,omitempty"`
	StartTradeDtTm     time.Time `datastore:"startTradeDtTm,noindex" json:"startTradeDtTm,omitempty"`
}

// Validate validates the TeamOpeningConfigPayload type instance.
func (ut *TeamOpeningConfigPayload) Validate() (err error) {

	return
}

// Team Description.
type teamPayload struct {
	// Team Win-Loss Record
	CurrentWinRecord *string `datastore:"currentWinRecord,noindex" json:"currentWinRecord,omitempty"`
	// Team Logo
	FullLogo *string `datastore:"fullLogo,noindex" json:"fullLogo,omitempty"`
	// Sport HomeTown ID
	HomeTownID *string `datastore:"homeTownId,noindex" json:"homeTownId,omitempty"`
	// Team Icon
	IconName *string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Team name
	Name *string `datastore:"name,noindex" json:"name,omitempty"`
	// Team Nickname
	ShortName *string `datastore:"shortName,noindex" json:"shortName,omitempty"`
	// Sport ID
	SportID *string `datastore:"sportId,noindex" json:"sportId,omitempty"`
}

// Validate validates the teamPayload type instance.
func (ut *teamPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.SportID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "sportId"))
	}
	if ut.HomeTownID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "homeTownId"))
	}
	if ut.ShortName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "shortName"))
	}
	if ut.CurrentWinRecord == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "currentWinRecord"))
	}
	if ut.IconName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "iconName"))
	}
	if ut.FullLogo == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "fullLogo"))
	}
	return
}

// Publicize creates TeamPayload from teamPayload
func (ut *teamPayload) Publicize() *TeamPayload {
	var pub TeamPayload
	if ut.CurrentWinRecord != nil {
		pub.CurrentWinRecord = *ut.CurrentWinRecord
	}
	if ut.FullLogo != nil {
		pub.FullLogo = *ut.FullLogo
	}
	if ut.HomeTownID != nil {
		pub.HomeTownID = *ut.HomeTownID
	}
	if ut.IconName != nil {
		pub.IconName = *ut.IconName
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.ShortName != nil {
		pub.ShortName = *ut.ShortName
	}
	if ut.SportID != nil {
		pub.SportID = *ut.SportID
	}
	return &pub
}

// Team Description.
type TeamPayload struct {
	// Team Win-Loss Record
	CurrentWinRecord string `datastore:"currentWinRecord,noindex" json:"currentWinRecord,omitempty"`
	// Team Logo
	FullLogo string `datastore:"fullLogo,noindex" json:"fullLogo,omitempty"`
	// Sport HomeTown ID
	HomeTownID string `datastore:"homeTownId,noindex" json:"homeTownId,omitempty"`
	// Team Icon
	IconName string `datastore:"iconName,noindex" json:"iconName,omitempty"`
	// Team name
	Name string `datastore:"name,noindex" json:"name,omitempty"`
	// Team Nickname
	ShortName string `datastore:"shortName,noindex" json:"shortName,omitempty"`
	// Sport ID
	SportID string `datastore:"sportId,noindex" json:"sportId,omitempty"`
}

// Validate validates the TeamPayload type instance.
func (ut *TeamPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.SportID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "sportId"))
	}
	if ut.HomeTownID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "homeTownId"))
	}
	if ut.ShortName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "shortName"))
	}
	if ut.CurrentWinRecord == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "currentWinRecord"))
	}
	if ut.IconName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "iconName"))
	}
	if ut.FullLogo == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "fullLogo"))
	}
	return
}
