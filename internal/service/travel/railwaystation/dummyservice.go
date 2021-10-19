package railwaystation

import "github.com/ozonmp/omp-bot/internal/model/travel"

type DummyRailwayStationService struct{}

func NewDummyRailwayStationService() *DummyRailwayStationService {
	return &DummyRailwayStationService{}
}

func (s *DummyRailwayStationService) Describe(stationID uint64) (*travel.RailwayStation, error) {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) List(cursor uint64, limit uint64) ([]travel.RailwayStation, error) {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Create(station travel.RailwayStation) (uint64, error) {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Update(stationID uint64, station travel.RailwayStation) error {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Remove(stationID uint64) (bool, error) {
	panic("Not implemented")
}
