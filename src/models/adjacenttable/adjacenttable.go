package adjacenttable

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/profile"
	"github.com/0B1t322/Magic-Circle/models/sector"
)

type Sector = sector.Sector
type Profile = profile.Profile

type AdjacentTable struct {
	ID          int     	`json:"id"`
	Sector      Sector  	`json:"sector"`
	Profle		Profile		`json:"profile"`
}

func AdjacentTableFromEnt(a *ent.AdjacentTable) AdjacentTable {
	return AdjacentTable{
		ID: a.ID,
		Sector: sector.NewSector(a.Edges.Sector),
		Profle: profile.ProfileFromEnt(a.Edges.Profile),
	}
}

func AdjacentTablesFromEnt(as []*ent.AdjacentTable) (slice []AdjacentTable) {
	for _, a := range as {
		slice = append(slice, AdjacentTableFromEnt(a))
	}
	return slice
}