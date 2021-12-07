package direction

import "github.com/0B1t322/Magic-Circle/ent"

type Direction struct {
	ID		int		`json:"id"`
	Name 	string	`json:"name"`
}


func DirectionFromEnt(d *ent.Direction) Direction {
	return Direction{
		ID: d.ID,
		Name: d.Name,
	}
}

func DirectionsFromEnt(ds []*ent.Direction) (slice []Direction) {
	for _, d := range ds {
		slice = append(slice, DirectionFromEnt(d))
	}
	return slice
}