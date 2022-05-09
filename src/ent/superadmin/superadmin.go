// Code generated by entc, DO NOT EDIT.

package superadmin

const (
	// Label holds the string label denoting the superadmin type in the database.
	Label = "super_admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLogin holds the string denoting the login field in the database.
	FieldLogin = "login"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// Table holds the table name of the superadmin in the database.
	Table = "SuperAdmin"
)

// Columns holds all SQL columns for superadmin fields.
var Columns = []string{
	FieldID,
	FieldLogin,
	FieldPassword,
	FieldEmail,
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