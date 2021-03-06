package app

import (
	"time"
)

// Fixture domain entity.
type Fixture struct {
	ID         uint64    `json:"id"`
	SeasonID   uint64    `json:"season_id"`
	RoundID    *uint64   `json:"round_id"`
	VenueID    *uint64   `json:"venue_id"`
	HomeTeamID uint64    `json:"home_team_id"`
	AwayTeamID uint64    `json:"away_team_id"`
	RefereeID  *uint64   `json:"referee_id"`
	Date       time.Time `json:"date"`
	Status     *string   `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// FixtureRepository provides an interface to persist Fixture domain struct objects to a storage engine.
type FixtureRepository interface {
	Insert(f *Fixture) error
	Update(f *Fixture) error
	Delete(id uint64) error
	ByID(id uint64) (*Fixture, error)
	ByTeamID(id uint64, query FixtureFilterQuery) ([]Fixture, error)
	Get(q FixtureRepositoryQuery) ([]Fixture, error)
	GetIDs(q FixtureRepositoryQuery) ([]uint64, error)
}

// FixtureRequester provides an interface allowing this application to request data from an external
// data provider. The requester implementation is responsible for creating the channel, filtering struct data into
// the channel before closing the channel once successful execution is complete.
type FixtureRequester interface {
	FixturesBySeasonIDs(ids []uint64) <-chan Fixture
}

type FixtureFilterQuery struct {
	DateAfter  *time.Time
	DateBefore *time.Time
	Limit      *uint64
	SeasonIDs  []uint64
	SortBy     *string
	Venue      *string
}

type FixtureRepositoryQuery struct {
	LeagueIDs        []uint64
	Filters          []FixtureStatFilter
	SeasonIDs        []uint64
	HomeTeamID       *uint64
	AwayTeamID       *uint64
	HomeTeamNameLike *string
	AwayTeamNameLike *string
	DateFrom         *time.Time
	DateTo           *time.Time
	Limit            *uint64
	SortBy           *string
}

type FixtureStatFilter struct {
	Type    string  `json:"type"`
	Team    string  `json:"team"`
	Metric  string  `json:"metric"`
	Measure string  `json:"measure"`
	Value   float32 `json:"value"`
	Venue   string  `json:"venue"`
	Games   uint8   `json:"game"`
}
