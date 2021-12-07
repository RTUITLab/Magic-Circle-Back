package profile

import "github.com/0B1t322/Magic-Circle/ent"

type Profile struct {
	ID 		int			`json:"id"`
	Name	string		`json:"name"`
}

func ProfileFromEnt(p *ent.Profile) Profile {
	return Profile{
		ID: p.ID,
		Name: p.Name,
	}
}

func ProfilesFromEnt(ps []*ent.Profile) (slice []Profile) {
	for _, p := range ps {
		slice = append(slice, ProfileFromEnt(p))
	}

	return slice
}