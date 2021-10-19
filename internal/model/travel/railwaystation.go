package travel

import "fmt"

type RailwayStation struct {
	Name     string
	Location string
}

func (s *RailwayStation) String() string {
	return fmt.Sprintf("%s (%s)", s.Name, s.Location)
}
