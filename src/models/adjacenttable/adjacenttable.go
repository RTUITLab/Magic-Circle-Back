package adjacenttable

import (
	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/models/sector"
	"github.com/0B1t322/Magic-Circle/models/variant"
)

type Sector = sector.Sector
type Variant = variant.Variant

type AdjacentTable struct {
	ID          int     `json:"id"`
	Sector      Sector  `json:"sector"`
	Variant     Variant `json:"variant"`
}

func AdjacentTableFromEnt(a *ent.AdjacentTable) AdjacentTable {
	return AdjacentTable{
		ID: a.ID,
		Sector: sector.NewSector(a.Edges.Sector),
		Variant: variant.VariantFromEnt(a.Edges.Variant),
	}
}

func AdjacentTablesFromEnt(as []*ent.AdjacentTable) (slice []AdjacentTable) {
	for _, a := range as {
		slice = append(slice, AdjacentTableFromEnt(a))
	}
	return slice
}