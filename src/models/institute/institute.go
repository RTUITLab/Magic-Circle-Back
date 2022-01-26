package institute

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/direction"
)

type Institute struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Directions []direction.Direction `json:"directions"`
}

func InstituteFromEnt(i *ent.Institute) Institute {
	return Institute{
		ID:   i.ID,
		Name: i.Name,
		Directions: direction.DirectionsFromEnt(i.Edges.Directions),
	}
}

func InstitutesFromEnt(is []*ent.Institute) (slice []Institute) {
	for _, i := range is {
		slice = append(slice, InstituteFromEnt(i))
	}

	return slice
}
