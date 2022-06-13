package sector

import (
	"github.com/0B1t322/Magic-Circle/ent"
)

type Profile struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	AdditionalDecription string `json:"additionalDescription"`
}
type Institute struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Directions []Direction `json:"directions"`
}
type Direction struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Profiles []Profile `json:"profiles"`
}

type AdditionalDescription struct {
	Institute Institute `json:"institute"`
}

func NewAdditionalDescription(a *ent.AdjacentTable) AdditionalDescription {
	return AdditionalDescription{
		Institute: Institute{
			ID:   a.Edges.Profile.Edges.Direction.Edges.Institute.ID,
			Name: a.Edges.Profile.Edges.Direction.Edges.Institute.Name,
			Directions: []Direction{
				{
					ID:   a.Edges.Profile.Edges.Direction.ID,
					Name: a.Edges.Profile.Edges.Direction.Name,
					Profiles: []Profile{
						{
							ID:                   a.Edges.Profile.ID,
							Name:                 a.Edges.Profile.Name,
							AdditionalDecription: a.AdditionalDescription,
						},
					},
				},
			},
		},
	}
}

func NewAdditionalDescriptions(as []*ent.AdjacentTable) (slice []AdditionalDescription) {
	for _, a := range as {
		slice = append(slice, NewAdditionalDescription(a))
	}
	return slice
}

// Return slice of unique instutes
func NewInstitues(ads []AdditionalDescription) (slice []Institute) {
	set := map[int]Institute{}
	for _, ad := range ads {
		if _, find := set[ad.Institute.ID]; !find {
			set[ad.Institute.ID] = ad.Institute
		}
	}

	for _, inst := range set {
		slice = append(slice, inst)
	}

	return slice
}

func FindDirectionsForInstitue(ads []AdditionalDescription, instId int) (slice []Direction) {
	set := map[int]Direction{}

	for _, ad := range ads {
		if ad.Institute.ID == instId {
			for _, dir := range ad.Institute.Directions {
				if _, find := set[dir.ID]; !find {
					set[dir.ID] = dir
				}
			}
		}
	}

	for _, dir := range set {
		slice = append(slice, dir)
	}

	return slice
}

func FindProfilesForDirection(ads []AdditionalDescription, instId, dirId int) (slice []Profile) {
	set := map[int]Profile{}

	for _, ad := range ads {
		if ad.Institute.ID == instId {
			for _, dir := range ad.Institute.Directions {
				if dirId == dir.ID {
					for _, prof := range dir.Profiles {
						if _, find := set[prof.ID]; !find {
							set[prof.ID] = prof
						}
					}
				}
			}
		}
	}

	for _, prof := range set {
		slice = append(slice, prof)
	}

	return slice
}

// func NewInstutesFromSector(s *ent.Sector) (insts []Institute) {
// 	set := map[int]Institute{}

// 	for _, a := range s.Edges.AdjacentTables {
// 		if inst, find := set[a.Edges.Profile.Edges.Direction.Edges.Institute.ID]; !find {
// 			set[a.Edges.Profile.Edges.Direction.Edges.Institute.ID] = Institute{
// 				ID:   a.Edges.Profile.Edges.Direction.Edges.Institute.ID,
// 				Name: a.Edges.Profile.Edges.Direction.Edges.Institute.Name,
// 				Directions: []Direction{
// 					{
// 						ID:   a.Edges.Profile.Edges.Direction.ID,
// 						Name: a.Edges.Profile.Edges.Direction.Name,
// 						Profiles: []Profile{
// 							{
// 								ID:                   a.Edges.Profile.ID,
// 								Name:                 a.Edges.Profile.Name,
// 								AdditionalDecription: a.AdditionalDescription,
// 							},
// 						},
// 					},
// 				},
// 			}
// 		} else if !find {
// 			// try insert direction

// 		}

// 	}
// }

type Sector struct {
	ID                    int                     `json:"id"`
	Coords                string                  `json:"coords"`
	Description           string                  `json:"description"`
	Institutes            []Institute             `json:"institutes,omitempty"`
}

// Institues not set
func NewSector(s *ent.Sector) Sector {
	return Sector{
		ID:          s.ID,
		Coords:      s.Coords,
		Description: s.Description,
		Institutes: NewInstitutesFromSector(s),
	}
}

func NewSectors(ss []*ent.Sector) (slice []Sector) {
	for _, s := range ss {
		slice = append(slice, NewSector(s))
	}
	return slice
}

func NewInstitutesFromSector(s *ent.Sector) []Institute {
	insts := map[int]*ent.Institute{}
	dirs :=  map[int]*ent.Direction{}
	profs :=  map[int]*ent.Profile{}

	for _, aj := range s.Edges.AdjacentTables {
		inst := aj.Edges.Profile.Edges.Direction.Edges.Institute
		dir := aj.Edges.Profile.Edges.Direction
		prof := aj.Edges.Profile
		{
			if _, find := insts[inst.ID]; !find {
				insts[inst.ID] = inst
			}

			if _, find := dirs[dir.ID]; !find {
				dirs[dir.ID] = dir
			}

			if _, find := profs[prof.ID]; !find {
				prof.Edges.AdjacentTables = append(prof.Edges.AdjacentTables, aj)
				profs[prof.ID] = prof
			}
		}
	}

	for _, prof := range profs {
		dir := dirs[prof.DirectionID]
		dir.Edges.Profile = append(dir.Edges.Profile, prof)
		dirs[prof.DirectionID] = dir
	}

	for _, dir := range dirs {
		inst := insts[dir.InstituteID]
		inst.Edges.Directions = append(inst.Edges.Directions, dir)
		insts[dir.InstituteID] = inst
	}

	var instSlice []*ent.Institute
	{
		for _, inst := range insts {
			instSlice = append(instSlice, inst)
		}
	}

	return NewInstitututes(instSlice)
}

func NewInstitututes(insts []*ent.Institute) (slice []Institute) {
	for _, i := range insts {
		slice = append(slice, NewInstitute(i))
	}

	return slice
}

func NewInstitute(inst *ent.Institute) Institute {
	return Institute{
		ID:         inst.ID,
		Name:       inst.Name,
		Directions: NewDirections(inst.Edges.Directions),
	}
}

func NewDirections(dirs []*ent.Direction) (slice []Direction) {
	for _, d := range dirs {
		slice = append(slice, NewDirection(d))
	}
	return slice
}

func NewDirection(dir *ent.Direction) Direction {
	return Direction{
		ID:       dir.ID,
		Name:     dir.Name,
		Profiles: NewProfiles(dir.Edges.Profile),
	}
}

func NewProfiles(profs []*ent.Profile) (slice []Profile) {
	for _, p := range profs {
		slice = append(slice, NewProfile(p))
	}

	return slice
}

func NewProfile(prof *ent.Profile) Profile {
	return Profile{
		ID:                   prof.ID,
		Name:                 prof.Name,
		AdditionalDecription: prof.Edges.AdjacentTables[0].AdditionalDescription,
	}
}