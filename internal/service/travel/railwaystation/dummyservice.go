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
		nextID:   1, // Reserve id=0 for invalid records
	}
	for i := 0; i < 15; i++ {
		s.Create(travel.RailwayStation{
			Name:     "Example " + strconv.Itoa(i+1),
			Location: "Some location " + strconv.Itoa(i+1),
		})
	}
	return s
}

func (s *DummyService) Describe(stationID uint64) (station *travel.RailwayStation, err error) {
	if stationID == 0 || stationID >= uint64(len(s.stations)) {
		err = fmt.Errorf("station with id %d doesn't exists", stationID)
		return
	}
	stationIdx := stationID - 1
	if s.stations[stationIdx].Deleted {
		err = fmt.Errorf("station with id %d was deleted", stationID)
		return
	}
	station = &s.stations[stationIdx].Station
	return
}

func (s *DummyService) List(cursor uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {
	if cursor > uint64(len(s.stations)) {
		cursor = uint64(len(s.stations)) + 1
	}
	currIdx := int(cursor) - 1
	if currIdx < 0 {
		currIdx = 0
	}
	for len(stationsArr) < int(limit) && currIdx < len(s.stations) {
		if !s.stations[currIdx].Deleted {
			stationsArr = append(stationsArr, s.stations[currIdx].Station)
		}
		currIdx++
	}
	return
}

func (s *DummyService) ListUntil(until uint64, limit uint64) (stationsArr []travel.RailwayStation, err error) {
	untilIdx := int(until) - 1
	if untilIdx > len(s.stations) {
		untilIdx = len(s.stations)
	}
	currIdx := untilIdx - 1
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
	if stationID == 0 || stationID >= uint64(len(s.stations)) {
		return fmt.Errorf("Station with id %d doesn't exists", stationID)
	}
	stationIdx := stationID - 1
	if s.stations[stationIdx].Deleted {
		return fmt.Errorf("Station with id %d was already deleted", stationID)
	}
	s.stations[stationIdx].Station = station
	return nil
}

func (s *DummyService) Remove(stationID uint64) (bool, error) {
	if stationID == 0 || stationID >= uint64(len(s.stations)) {
		return false, fmt.Errorf("Station with id %d doesn't exists", stationID)
	}
	stationIdx := stationID - 1
	if s.stations[stationIdx].Deleted {
		return false, fmt.Errorf("Station with id %d was already deleted", stationID)
	}
	s.stations[stationIdx].Deleted = true
	return true, nil
}
