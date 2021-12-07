package sector

import "github.com/0B1t322/Magic-Circle/ent"

type Sector struct {
	ID          int    `json:"id"`
	Coords      string `json:"coords"`
	Description string `json:"description"`
}

func NewSector(s *ent.Sector) Sector {
	return Sector{
		ID:          s.ID,
		Coords:      s.Coords,
		Description: s.Description,
	}
}

func NewSectors(ss []*ent.Sector) (slice []Sector) {
	for _, s := range ss {
		slice = append(slice, NewSector(s))
	}
	return slice
}