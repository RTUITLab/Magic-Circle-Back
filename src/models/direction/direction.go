package direction

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/profile"
)

type Direction struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Profiles []profile.Profile `json:"profiles"`
}

func DirectionFromEnt(d *ent.Direction) Direction {
	return Direction{
		ID:   d.ID,
		Name: d.Name,
		Profiles: profile.ProfilesFromEnt(d.Edges.Profile),
	}
}

func DirectionsFromEnt(ds []*ent.Direction) (slice []Direction) {
	for _, d := range ds {
		slice = append(slice, DirectionFromEnt(d))
	}
	return slice
}
