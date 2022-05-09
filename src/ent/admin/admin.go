// Code generated by entc, DO NOT EDIT.

package admin

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldInstituteID holds the string denoting the institute_id field in the database.
	FieldInstituteID = "institute_id"
	// EdgeInstitute holds the string denoting the institute edge name in mutations.
	EdgeInstitute = "Institute"
	// Table holds the table name of the admin in the database.
	Table = "Admin"
	// InstituteTable is the table that holds the Institute relation/edge.
	InstituteTable = "Admin"
	// InstituteInverseTable is the table name for the Institute entity.
	// It exists in this package in order to avoid circular dependency with the "institute" package.
	InstituteInverseTable = "Institute"
	// InstituteColumn is the table column denoting the Institute relation/edge.
	InstituteColumn = "institute_id"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldPassword,
	FieldInstituteID,
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