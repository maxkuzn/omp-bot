package railwaystation

import (
	"fmt"
	"strconv"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type stationInfo struct {
	Deleted bool
	Station travel.RailwayStation
}

type DummyService struct {
	stations []stationInfo
	nextID   uint64
}

func NewDummyService() *DummyService {
	stations := []stationInfo{}
	s := &DummyService{
		stations: stations,
		nextID:   0,
	}
	for i := 0; i < 15; i++ {
		s.Create(travel.RailwayStation{
			Name:     "Example " + strconv.Itoa(i),
			Location: "Some location " + strconv.Itoa(i),
		})
	}
	return s
}

func (s *DummyService) Describe(stationID uint64) (station *travel.RailwayStation, err error) {
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

func (s *DummyService) List(cursor uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {
	if cursor > uint64(len(s.stations)) {
		cursor = uint64(len(s.stations))
	}
	currIdx := int(cursor)
	for len(stationsArr) < int(limit) && currIdx < len(s.stations) {
		if !s.stations[currIdx].Deleted {
			stationsArr = append(stationsArr, s.stations[currIdx].Station)
		}
		currIdx++
	}
	return
}

func (s *DummyService) ListUntil(until uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {
	if until > uint64(len(s.stations)) {
		until = uint64(len(s.stations))
	}
	currIdx := int(until) - 1
	stationsArr = make([]travel.RailwayStation, limit)
	currLen := 0
	for currLen < int(limit) && currIdx >= 0 {
		if !s.stations[currIdx].Deleted {
			stationsArr[int(limit)-currLen-1] = s.stations[currIdx].Station
			currLen++
		}
		currIdx--
	}
	stationsArr = stationsArr[int(limit)-currLen:]
	return
}

func (s *DummyService) Create(station travel.RailwayStation) (uint64, error) {
	station.ID = s.nextID
	s.nextID++

	info := stationInfo{
		Deleted: false,
		Station: station,
	}

	s.stations = append(s.stations, info)
	return station.ID, nil
}

func (s *DummyService) Update(stationID uint64, station travel.RailwayStation) error {
	if stationID >= uint64(len(s.stations)) {
		return fmt.Errorf("Station with id %d doesn't exists", stationID)
	}
	if s.stations[stationID].Deleted {
		return fmt.Errorf("Station with id %d was already deleted", stationID)
	}
	s.stations[stationID].Station = station
	return nil
}

func (s *DummyService) Remove(stationID uint64) (bool, error) {
	if stationID >= uint64(len(s.stations)) {
		return false, fmt.Errorf("Station with id %d doesn't exists", stationID)
	}
	if s.stations[stationID].Deleted {
		return false, fmt.Errorf("Station with id %d was already deleted", stationID)
	}
	s.stations[stationID].Deleted = true
	return true, nil
}
