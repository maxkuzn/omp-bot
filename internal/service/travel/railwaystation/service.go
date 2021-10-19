package railwaystation

import "github.com/ozonmp/omp-bot/internal/model/travel"

type Service interface {
	Describe(stationID uint64) (*travel.RailwayStation, error)
	List(cursor uint64, limit uint64) ([]travel.RailwayStation, error)
	Create(travel.RailwayStation) (uint64, error)
	Update(stationID uint64, station travel.RailwayStation) error
	Remove(stationID uint64) (bool, error)
}
