package mock

import (
	"github.com/statistico/statistico-data/internal/app"
	"github.com/stretchr/testify/mock"
)

type TeamRepository struct {
	mock.Mock
}

func (m TeamRepository) Insert(c *app.Team) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m TeamRepository) Update(c *app.Team) error {
	args := m.Called(&c)
	return args.Error(0)
}

func (m TeamRepository) ByID(id int64) (*app.Team, error) {
	args := m.Called(id)
	c := args.Get(0).(*app.Team)
	return c, args.Error(1)
}

type TeamRequester struct {
	mock.Mock
}

func (t TeamRequester) TeamsBySeasonIDs(seasonIDs []int64) <-chan *app.Team {
	args := t.Called(seasonIDs)
	return args.Get(0).(chan *app.Team)
}