package railwaystation

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type DummyRailwayStationService struct {
	stations map[uint64]travel.RailwayStation
	nextID   uint64
}

func NewDummyRailwayStationService() *DummyRailwayStationService {
	stations := make(map[uint64]travel.RailwayStation)
	stations[0] = travel.RailwayStation{
		ID:       0,
		Name:     "Example",
		Location: "Some location",
	}
	return &DummyRailwayStationService{
		stations: stations,
		nextID:   1,
	}
}

func (s *DummyRailwayStationService) Describe(stationID uint64) (stationPtr *travel.RailwayStation, err error) {
	station, ok := s.stations[stationID]
	if !ok {
		err = fmt.Errorf("Station with id %d doesn't exists", stationID)
		return
	}
	stationPtr = &station
	return
}

func (s *DummyRailwayStationService) List(cursor uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Create(station travel.RailwayStation) (uint64, error) {
	station.ID = s.nextID
	s.nextID++

	s.stations[station.ID] = station
	return station.ID, nil
}

func (s *DummyRailwayStationService) Update(stationID uint64, station travel.RailwayStation) error {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Remove(stationID uint64) (bool, error) {
	panic("Not implemented")
}
