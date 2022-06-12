package sector

import (
	"github.com/0B1t322/Magic-Circle/ent"
)

type InstituteCreator struct {
	// Key is profileId
	DescriptionMap map[int]string

}

func NewInstituteCreator(descriptionMap map[int]string) *InstituteCreator {
	return &InstituteCreator{
		DescriptionMap: descriptionMap,
	}
}

func (i *InstituteCreator) NewInstitututes(insts []*ent.Institute) (slice []Institute) {
	for _, inst := range insts {
		slice = append(slice, i.NewInstitute(inst))
	}

	return slice
}

func (i *InstituteCreator) NewInstitute(inst *ent.Institute) Institute {
	return Institute{
		ID:         inst.ID,
		Name:       inst.Name,
		Directions: i.NewDirections(inst.Edges.Directions),
	}
}

func (i *InstituteCreator) NewDirections(dirs []*ent.Direction) (slice []Direction) {
	for _, d := range dirs {
		slice = append(slice, i.NewDirection(d))
	}
	return slice
}

func (i *InstituteCreator) NewDirection(dir *ent.Direction) Direction {
	return Direction{
		ID:       dir.ID,
		Name:     dir.Name,
		Profiles: i.NewProfiles(dir.Edges.Profile),
	}
}

func (i *InstituteCreator) NewProfiles(profs []*ent.Profile) (slice []Profile) {
	for _, p := range profs {
		slice = append(slice, i.NewProfile(p))
	}

	return slice
}

func (i *InstituteCreator) NewProfile(prof *ent.Profile) Profile {
	return Profile{
		ID:                   prof.ID,
		Name:                 prof.Name,
		AdditionalDecription: i.DescriptionMap[prof.ID],
	}
}