package models

// Database is the representation of a relational database managed service
// Can be RDS instance for AWS or Azure Database instance for Azure.
type Database struct {

	// Engine is the type of the relational database. For instance
	// postgres or mariadb
	Engine string `json:"engine"`

	// Version of the engine
	Version string `json:"version"`

	// Size of the underlying compute units used for
	// this database
	Size string `json:"size"`
}
