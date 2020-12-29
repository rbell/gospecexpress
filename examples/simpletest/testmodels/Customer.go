package testmodels

import "time"

// Customer is a sample customer model for purposes of validation
type Customer struct {
	FirstName      string
	middleName     string
	LastName       string
	Age            int
	MemberSince    time.Time
	MemberExpireAt time.Time
}
