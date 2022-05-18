package sector

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/profile"
)

type Profile = profile.Profile

type Direction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func DirectionFromEnt(d *ent.Direction) Direction {
	return Direction{
		ID:   d.ID,
		Name: d.Name,
	}
}

type Institute struct {
	ID   int    `json:"id"`
	Name string `json:"name"` 
}

func InstituteFromEnt(i *ent.Institute) Institute {
	return Institute{
		ID:   i.ID,
		Name: i.Name,
	}
}

type AdditionalDescription struct {
	Profile              Profile             `json:"profile"`
	Direction            Direction `json:"direction"`
	Institute            Institute `json:"institute"`
	AdditionalDecription string              `json:"additionalDescription"`
}

func NewAdditionalDescription(a *ent.AdjacentTable) AdditionalDescription {
	return AdditionalDescription{
		Profile:              profile.ProfileFromEnt(a.Edges.Profile),
		Direction:            DirectionFromEnt(a.Edges.Profile.Edges.Direction),
		Institute:            InstituteFromEnt(a.Edges.Profile.Edges.Direction.Edges.Institute),
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
