package railwaystation

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type stationInfo struct {
	Deleted bool
	Station travel.RailwayStation
}

type DummyRailwayStationService struct {
	stations []stationInfo
	nextID   uint64
}

func NewDummyRailwayStationService() *DummyRailwayStationService {
	info := stationInfo{
		Deleted: false,
		Station: travel.RailwayStation{
			ID:       0,
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
	currIdx := int(cursor)
	for len(stationsArr) < int(limit) && currIdx < len(s.stations) {
		if !s.stations[currIdx].Deleted {
			stationsArr = append(stationsArr, s.stations[currIdx].Station)
		}
		currIdx++
	}
	return
}

func (s *DummyRailwayStationService) Create(station travel.RailwayStation) (uint64, error) {
	station.ID = s.nextID
	s.nextID++

	info := stationInfo{
		Deleted: false,
		Station: station,
	}

	s.stations = append(s.stations, info)
	return station.ID, nil
}

func (s *DummyRailwayStationService) Update(stationID uint64, station travel.RailwayStation) error {
	panic("Not implemented")
}

func (s *DummyRailwayStationService) Remove(stationID uint64) (bool, error) {
	if stationID >= uint64(len(s.stations)) {
		return false, fmt.Errorf("Station with id %d doesn't exists", stationID)
	}
	if s.stations[stationID].Deleted {
		return false, fmt.Errorf("Station with id %d was already deleted", stationID)
	}
	s.stations[stationID].Deleted = true
	return true, nil
}
