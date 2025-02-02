// Code generated by ent, DO NOT EDIT.

package withmodifiedfield

const (
	// Label holds the string label denoting the withmodifiedfield type in the database.
	Label = "with_modified_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the withmodifiedfield in the database.
	Table = "with_modified_fields"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "with_modified_fields"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "with_modified_field_owner"
)

// Columns holds all SQL columns for withmodifiedfield fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "with_modified_fields"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"with_modified_field_owner",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
