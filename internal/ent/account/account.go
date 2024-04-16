// Code generated by ent, DO NOT EDIT.

package account

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldDisableTime holds the string denoting the disable_time field in the database.
	FieldDisableTime = "disable_time"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "profiles"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "account_profile"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldDisableTime,
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

var (
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
