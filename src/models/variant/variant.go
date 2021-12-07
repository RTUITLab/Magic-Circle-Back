package variant

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/direction"
	"github.com/0B1t322/Magic-Circle/models/institute"
	"github.com/0B1t322/Magic-Circle/models/profile"
)


type Direction = direction.Direction
type Institute = institute.Institute
type Profile = profile.Profile

type Variant struct {
	ID			int			`json:"id"`
	Direction	Direction	`json:"direction"`
	Institute	Institute	`json:"institute"`
	Profile		Profile		`json:"profile"`
}

// Should contains all edges
func VariantFromEnt(v *ent.Variant) Variant {
	return Variant{
		ID: v.ID,
		Direction: direction.DirectionFromEnt(v.Edges.Direction),
		Institute: institute.InstituteFromEnt(v.Edges.Insitute),
		Profile: profile.ProfileFromEnt(v.Edges.Profile),
	}
}

func VariantsFromEnt(vs []*ent.Variant) (slice []Variant) {
	for _, v := range vs {
		slice = append(slice, VariantFromEnt(v))
	}
	return slice
}