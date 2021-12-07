// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/Magic-Circle/ent/sector"
)

// Sector is the model entity for the Sector schema.
type Sector struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Coords holds the value of the "coords" field.
	Coords string `json:"coords,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SectorQuery when eager-loading is set.
	Edges SectorEdges `json:"edges"`
}

// SectorEdges holds the relations/edges for other nodes in the graph.
type SectorEdges struct {
	// AdjacentTables holds the value of the AdjacentTables edge.
	AdjacentTables []*AdjacentTable `json:"AdjacentTables,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AdjacentTablesOrErr returns the AdjacentTables value or an error if the edge
// was not loaded in eager-loading.
func (e SectorEdges) AdjacentTablesOrErr() ([]*AdjacentTable, error) {
	if e.loadedTypes[0] {
		return e.AdjacentTables, nil
	}
	return nil, &NotLoadedError{edge: "AdjacentTables"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sector) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sector.FieldID:
			values[i] = new(sql.NullInt64)
		case sector.FieldCoords, sector.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sector", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sector fields.
func (s *Sector) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sector.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sector.FieldCoords:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coords", values[i])
			} else if value.Valid {
				s.Coords = value.String
			}
		case sector.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		}
	}
	return nil
}

// QueryAdjacentTables queries the "AdjacentTables" edge of the Sector entity.
func (s *Sector) QueryAdjacentTables() *AdjacentTableQuery {
	return (&SectorClient{config: s.config}).QueryAdjacentTables(s)
}

// Update returns a builder for updating this Sector.
// Note that you need to call Sector.Unwrap() before calling this method if this Sector
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sector) Update() *SectorUpdateOne {
	return (&SectorClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sector entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sector) Unwrap() *Sector {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sector is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sector) String() string {
	var builder strings.Builder
	builder.WriteString("Sector(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", coords=")
	builder.WriteString(s.Coords)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Sectors is a parsable slice of Sector.
type Sectors []*Sector

func (s Sectors) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
