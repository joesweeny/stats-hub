package postgres_test

import (
	"github.com/statistico/statistico-football-data/internal/app"
	"github.com/statistico/statistico-football-data/internal/app/postgres"
	"github.com/statistico/statistico-football-data/internal/app/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSquadRepository_Insert(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_squad")
	repo := postgres.NewSquadRepository(conn, test.Clock)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i < 4; i++ {
			c := newSquad(uint64(i), uint64(i+1))

			if err := repo.Insert(c); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from sportmonks_squad")

			var count int

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, i, count)
		}
	})
}

func TestSquadRepository_BySeasonAndTeam(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_squad")
	repo := postgres.NewSquadRepository(conn, test.Clock)

	t.Run("squad can be retrieved by season and team IDs", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newSquad(45, 986)

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		r, err := repo.BySeasonAndTeam(45, 986)

		if err != nil {
			t.Errorf("Error when retrieving a record from the database: %s", err.Error())
		}

		a := assert.New(t)
		a.Equal(uint64(45), r.SeasonID)
		a.Equal(uint64(986), m.TeamID)
		a.Equal([]uint64{34, 57, 89}, m.PlayerIDs)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns error if record does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		if _, err := repo.BySeasonAndTeam(99, 2435870); err == nil {
			t.Fatalf("Test failed, expected %v, got nil", err)
		}
	})
}

func TestSquadRepository_Update(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_squad")
	repo := postgres.NewSquadRepository(conn, test.Clock)

	t.Run("modifies existing squad", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newSquad(25, 62)

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		m.PlayerIDs = []uint64{uint64(432), uint64(567), uint64(2), uint64(87095)}
		m.UpdatedAt = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)

		if err := repo.Update(m); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.BySeasonAndTeam(25, 62)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		a := assert.New(t)
		a.Equal(uint64(25), r.SeasonID)
		a.Equal(uint64(62), m.TeamID)
		a.Equal([]uint64{432, 567, 2, 87095}, m.PlayerIDs)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns an error if player does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newSquad(146, 200)

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}
	})
}

func TestSquadRepository_All(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_squad")
	repo := postgres.NewSquadRepository(conn, test.Clock)

	t.Run("returns all squad records from the database", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		squads := []app.Squad{
			*newSquad(9, 25067),
			*newSquad(99, 98),
			*newSquad(301, 2),
			*newSquad(23, 6),
			*newSquad(39, 1902),
		}

		for _, squad := range squads {
			err := repo.Insert(&squad)

			if err != nil {
				t.Fatalf("Test failed, expected nil, got %v", err)
			}
		}

		all, err := repo.All()

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}

		a := assert.New(t)

		a.Equal(5, len(all))
		a.Equal(all[:1], squads[0:1])
		a.Equal(all[1:2], squads[3:4])
		a.Equal(all[2:3], squads[4:5])
		a.Equal(all[3:4], squads[1:2])
		a.Equal(all[4:], squads[2:3])
	})
}

func TestSquadRepository_CurrentSeason(t *testing.T) {
	conn, cleanUpSquad := test.GetConnection(t, "sportmonks_squad")
	repo := postgres.NewSquadRepository(conn, test.Clock)
	conn, cleanUpSeason := test.GetConnection(t, "sportmonks_season")
	seasonRepo := postgres.NewSeasonRepository(conn, test.Clock)

	t.Run("returns squads only for current season", func(t *testing.T) {
		t.Helper()
		defer cleanUpSquad()
		defer cleanUpSeason()

		seasons := []*app.Season{
			newSeason(39, 560, "2018-2019", true),
			newSeason(99, 560, "2018-2019", false),
			newSeason(4502, 560, "2018-2019", false),
			newSeason(23, 560, "2018-2019", true),
		}

		for _, s := range seasons {
			err := seasonRepo.Insert(s)

			if err != nil {
				t.Fatalf("Test failed, expected nil, got %v", err)
			}
		}

		squads := []app.Squad{
			*newSquad(39, 25067),
			*newSquad(99, 98),
			*newSquad(301, 2),
			*newSquad(23, 6),
			*newSquad(39, 1902),
		}

		for _, squad := range squads {
			err := repo.Insert(&squad)

			if err != nil {
				t.Fatalf("Test failed, expected nil, got %v", err)
			}
		}

		current, err := repo.CurrentSeason()

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}

		a := assert.New(t)

		a.Equal(3, len(current))
		a.Equal(current[:1], squads[3:4])
		a.Equal(current[1:2], squads[4:])
		a.Equal(current[2:3], squads[0:1])
	})
}

func newSquad(season, team uint64) *app.Squad {
	return &app.Squad{
		SeasonID:  season,
		TeamID:    team,
		PlayerIDs: []uint64{34, 57, 89},
		CreatedAt: time.Unix(1547465100, 0),
		UpdatedAt: time.Unix(1547465100, 0),
	}
}
