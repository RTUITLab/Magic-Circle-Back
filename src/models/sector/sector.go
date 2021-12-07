package sector

import "github.com/0B1t322/Magic-Circle/ent"

type Sector struct {
	ID		int		`json:"int"`
	Coords	string	`json:"coords"`
}

func NewSector(s *ent.Sector) Sector {
	return Sector{
		ID: s.ID,
		Coords: s.Coords,
	}
}