package institute

import "github.com/0B1t322/Magic-Circle/ent"

type Institute struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
}

func InstituteFromEnt(i *ent.Institute) Institute {
	return Institute{
		ID: i.ID,
		Name: i.Name,
	}
}

func InstitutesFromEnt(is []*ent.Institute) (slice []Institute) {
	for _, i := range is {
		slice = append(slice, InstituteFromEnt(i))
	}

	return slice
}