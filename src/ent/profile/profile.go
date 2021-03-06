// Code generated by entc, DO NOT EDIT.

package profile

const (
	// Label holds the string label denoting the profile type in the database.
	Label = "profile"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDirectionID holds the string denoting the direction_id field in the database.
	FieldDirectionID = "direction_id"
	// EdgeDirection holds the string denoting the direction edge name in mutations.
	EdgeDirection = "Direction"
	// EdgeAdjacentTables holds the string denoting the adjacenttables edge name in mutations.
	EdgeAdjacentTables = "AdjacentTables"
	// Table holds the table name of the profile in the database.
	Table = "Profile"
	// DirectionTable is the table that holds the Direction relation/edge.
	DirectionTable = "Profile"
	// DirectionInverseTable is the table name for the Direction entity.
	// It exists in this package in order to avoid circular dependency with the "direction" package.
	DirectionInverseTable = "Direction"
	// DirectionColumn is the table column denoting the Direction relation/edge.
	DirectionColumn = "direction_id"
	// AdjacentTablesTable is the table that holds the AdjacentTables relation/edge.
	AdjacentTablesTable = "AdjacentTable"
	// AdjacentTablesInverseTable is the table name for the AdjacentTable entity.
	// It exists in this package in order to avoid circular dependency with the "adjacenttable" package.
	AdjacentTablesInverseTable = "AdjacentTable"
	// AdjacentTablesColumn is the table column denoting the AdjacentTables relation/edge.
	AdjacentTablesColumn = "profile_id"
)

// Columns holds all SQL columns for profile fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDirectionID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
