package sector

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/direction"
	"github.com/0B1t322/Magic-Circle/models/institute"
	"github.com/0B1t322/Magic-Circle/models/profile"
)

type Profile = profile.Profile
type Institute = institute.Institute
type Direction = direction.Direction

type AdditionalDescription struct {
	Institute            Institute `json:"institute"`
	AdditionalDecription string    `json:"additionalDescription"`
}

func NewAdditionalDescription(a *ent.AdjacentTable) AdditionalDescription {
	return AdditionalDescription{
		Institute:            Institute{
			ID: a.Edges.Profile.Edges.Direction.Edges.Institute.ID,
			Name: a.Edges.Profile.Edges.Direction.Edges.Institute.Name,
			Directions: []Direction{
				{
					ID:       a.Edges.Profile.Edges.Direction.ID,
					Name:     a.Edges.Profile.Edges.Direction.Name,
					Profiles: []Profile{
						{
							ID:   a.Edges.Profile.ID,
							Name: a.Edges.Profile.Name,
						},
					},
				},
			},
		},
		AdditionalDecription: a.AdditionalDescription,
	}
}

func NewAdditionalDescriptions(as []*ent.AdjacentTable) (slice []AdditionalDescription) {
	for _, a := range as {
		slice = append(slice, NewAdditionalDescription(a))
	}
	return slice
}

type Sector struct {
	ID                     int                     `json:"id"`
	Coords                 string                  `json:"coords"`
	Description            string                  `json:"description"`
	AdditionalDescriptions []AdditionalDescription `json:"additionalDescriptions"`
}

func NewSector(s *ent.Sector) Sector {
	return Sector{
		ID:                     s.ID,
		Coords:                 s.Coords,
		Description:            s.Description,
		AdditionalDescriptions: NewAdditionalDescriptions(s.Edges.AdjacentTables),
	}
}

func NewSectors(ss []*ent.Sector) (slice []Sector) {
	for _, s := range ss {
		slice = append(slice, NewSector(s))
	}
	return slice
}
