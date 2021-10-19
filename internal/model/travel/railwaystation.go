package travel

import "fmt"

type RailwayStation struct {
	ID       uint64
	Name     string
	Location string
}

func (s *RailwayStation) String() string {
	return fmt.Sprintf("[%d] %s (%s)", s.ID, s.Name, s.Location)
}
