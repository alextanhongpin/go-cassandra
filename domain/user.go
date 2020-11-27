package domain

import (
	"github.com/gocql/gocql"
)

type User struct {
	ID        gocql.UUID
	FirstName string
	LastName  string
	Email     string
	City      string
	Age       int
}
