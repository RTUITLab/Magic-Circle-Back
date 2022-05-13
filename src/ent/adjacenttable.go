// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/0B1t322/Magic-Circle/ent/adjacenttable"
	"github.com/0B1t322/Magic-Circle/ent/profile"
	"github.com/0B1t322/Magic-Circle/ent/sector"
)

// AdjacentTable is the model entity for the AdjacentTable schema.
type AdjacentTable struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SectorID holds the value of the "sector_id" field.
	SectorID int `json:"sector_id,omitempty"`
	// ProfileID holds the value of the "profile_id" field.
	ProfileID int `json:"profile_id,omitempty"`
	// AdditionalDescription holds the value of the "additionalDescription" field.
	AdditionalDescription string `json:"additionalDescription,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdjacentTableQuery when eager-loading is set.
	Edges AdjacentTableEdges `json:"edges"`
}

// AdjacentTableEdges holds the relations/edges for other nodes in the graph.
type AdjacentTableEdges struct {
	// Profile holds the value of the Profile edge.
	Profile *Profile `json:"Profile,omitempty"`
	// Sector holds the value of the Sector edge.
	Sector *Sector `json:"Sector,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdjacentTableEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Profile == nil {
			// The edge Profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "Profile"}
}

// SectorOrErr returns the Sector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdjacentTableEdges) SectorOrErr() (*Sector, error) {
	if e.loadedTypes[1] {
		if e.Sector == nil {
			// The edge Sector was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sector.Label}
		}
		return e.Sector, nil
	}
	return nil, &NotLoadedError{edge: "Sector"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdjacentTable) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case adjacenttable.FieldID, adjacenttable.FieldSectorID, adjacenttable.FieldProfileID:
			values[i] = new(sql.NullInt64)
		case adjacenttable.FieldAdditionalDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdjacentTable", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdjacentTable fields.
func (at *AdjacentTable) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adjacenttable.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			at.ID = int(value.Int64)
		case adjacenttable.FieldSectorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sector_id", values[i])
			} else if value.Valid {
				at.SectorID = int(value.Int64)
			}
		case adjacenttable.FieldProfileID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field profile_id", values[i])
			} else if value.Valid {
				at.ProfileID = int(value.Int64)
			}
		case adjacenttable.FieldAdditionalDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field additionalDescription", values[i])
			} else if value.Valid {
				at.AdditionalDescription = value.String
			}
		}
	}
	return nil
}

// QueryProfile queries the "Profile" edge of the AdjacentTable entity.
func (at *AdjacentTable) QueryProfile() *ProfileQuery {
	return (&AdjacentTableClient{config: at.config}).QueryProfile(at)
}

// QuerySector queries the "Sector" edge of the AdjacentTable entity.
func (at *AdjacentTable) QuerySector() *SectorQuery {
	return (&AdjacentTableClient{config: at.config}).QuerySector(at)
}

// Update returns a builder for updating this AdjacentTable.
// Note that you need to call AdjacentTable.Unwrap() before calling this method if this AdjacentTable
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AdjacentTable) Update() *AdjacentTableUpdateOne {
	return (&AdjacentTableClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the AdjacentTable entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AdjacentTable) Unwrap() *AdjacentTable {
	tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdjacentTable is not a transactional entity")
	}
	at.config.driver = tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AdjacentTable) String() string {
	var builder strings.Builder
	builder.WriteString("AdjacentTable(")
	builder.WriteString(fmt.Sprintf("id=%v", at.ID))
	builder.WriteString(", sector_id=")
	builder.WriteString(fmt.Sprintf("%v", at.SectorID))
	builder.WriteString(", profile_id=")
	builder.WriteString(fmt.Sprintf("%v", at.ProfileID))
	builder.WriteString(", additionalDescription=")
	builder.WriteString(at.AdditionalDescription)
	builder.WriteByte(')')
	return builder.String()
}

// AdjacentTables is a parsable slice of AdjacentTable.
type AdjacentTables []*AdjacentTable

func (at AdjacentTables) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}
