package utils

import (
	"github.com/0B1t322/Magic-Circle/ent/direction"
	"github.com/0B1t322/Magic-Circle/ent/institute"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
	"github.com/0B1t322/Magic-Circle/ent/profile"
)

func PredicatesByDirectionsNames(dirs []string) (preds []predicate.Direction) {
	for _, dir := range dirs {
		preds = append(
			preds, 
			direction.Name(dir),
		)
	}

	return preds
}

func PredicatesByInstititeNames(insts []string) (preds []predicate.Institute) {
	for _, inst := range insts {
		preds = append(
			preds, 
			institute.Name(inst),
		)
	}
	return preds
}

func PredicatesByProfilesName(profs []string) (preds []predicate.Profile) {
	for _, prof := range profs {
		preds = append(
			preds, 
			profile.Name(prof),
		)
	}
	return preds
}