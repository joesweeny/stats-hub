package fixture

import (
	"github.com/joesweeny/sportmonks-go-client"
	"github.com/joesweeny/statshub/internal/model"
	"github.com/joesweeny/statshub/internal/season"
	"log"
	"sync"
)

var waitGroup sync.WaitGroup

type Service struct {
	Repository
	SeasonRepo season.Repository
	Factory
	Client *sportmonks.Client
	Logger *log.Logger
}

func (s Service) Process() error {
	ids, err := s.SeasonRepo.Ids()

	if err != nil {
		return err
	}

	return s.callClient(ids)
}

func (s Service) CurrentSeason() error {
	ids, err := s.SeasonRepo.CurrentSeasonIds()

	if err != nil {
		return err
	}

	return s.callClient(ids)
}

func (s Service) callClient(ids []int) error {
	q := []string{"fixtures"}

	for _, id := range ids {
		go func(id int) {
			waitGroup.Add(1)
			res, err := s.Client.SeasonById(id, q)

			s.handleSeason(res.Data)

			if err != nil {
				log.Printf("Error when calling client '%s", err.Error())
			}
			defer waitGroup.Done()
		}(id)
	}

	waitGroup.Wait()

	return nil
}

func (s Service) handleSeason(m sportmonks.Season) {
	for _, fixture := range m.Fixtures.Data {
		waitGroup.Add(1)

		go func(fixture sportmonks.Fixture) {
			s.persistFixture(&fixture)
			defer waitGroup.Done()
		}(fixture)
	}
}

func (s Service) persistFixture(m *sportmonks.Fixture) {
	fixture, err := s.GetById(m.ID)

	if err != nil && (model.Fixture{}) == *fixture {
		created := s.createFixture(m)

		if err := s.Insert(created); err != nil {
			log.Printf("Error '%s' occurred when inserting Fixture struct: %+v\n,", err.Error(), created)
		}

		return
	}

	updated := s.updateFixture(m, fixture)

	if err := s.Update(updated); err != nil {
		log.Printf("Error '%s' occurred when updating Fixture struct: %+v\n,", err.Error(), updated)
	}

	return
}
