package railwaystation

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type stationInfo struct {
	ID      uint64
	Deleted bool
	Station travel.RailwayStation
}

type DummyRailwayStationService struct {
	stations []stationInfo
	nextID   uint64
}

func NewDummyRailwayStationService() *DummyRailwayStationService {
	info := stationInfo{
		ID:      0,
		Deleted: false,
		Station: travel.RailwayStation{
			Name:     "Example",
			Location: "Some location",
		},
	}
	stations := []stationInfo{info}
	return &DummyRailwayStationService{
		stations: stations,
		nextID:   1,
	}
}

func (s *DummyRailwayStationService) Describe(stationID uint64) (station *travel.RailwayStation, err error) {
	if stationID >= uint64(len(s.stations)) {
		err = fmt.Errorf("Station with id %d doesn't exists", stationID)
		return
	}
	if s.stations[stationID].Deleted {
		err = fmt.Errorf("Station with id %d was deleted", stationID)
		return
	}
	station = &s.stations[stationID].Station
	return
}

func (s *DummyRailwayStationService) List(cursor uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {

	panic("Not implemented")
}

func (s *DummyRailwayStationService) Create(station travel.RailwayStation) (uint64, error) {
	station.ID = s.nextID
	info := stationInfo{
		ID:      s.nextID,
		Deleted: false,
		Station: station,
	}
	s.nextID++

	s.stations = append(s.stations, info)
	return info.ID, nil
}

func (s *DummyRailwayStationService) Update(stationID uint64, station travel.RailwayStation) error {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Remove(stationID uint64) (bool, error) {
	panic("Not implemented")
}
