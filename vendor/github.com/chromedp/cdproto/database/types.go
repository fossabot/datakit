package database

// Code generated by cdproto-gen. DO NOT EDIT.

// ID unique identifier of Database object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Database#type-DatabaseId
type ID string

// String returns the ID as string value.
func (t ID) String() string {
	return string(t)
}

// Database database object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Database#type-Database
type Database struct {
	ID      ID     `json:"id"`      // Database ID.
	Domain  string `json:"domain"`  // Database domain.
	Name    string `json:"name"`    // Database name.
	Version string `json:"version"` // Database version.
}

// Error database error.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Database#type-Error
type Error struct {
	Message string `json:"message"` // Error message.
	Code    int64  `json:"code"`    // Error code.
}
