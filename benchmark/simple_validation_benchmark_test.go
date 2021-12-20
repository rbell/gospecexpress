package benchmark

import (
	"fmt"
	"testing"
	"time"

	"github.com/rbell/gospecexpress"
	"github.com/rbell/gospecexpress/catalog"
)

// ClubMember is a sample customer model for purposes of validation benchmark
// Field names can be tagged with a user readable field name used when referencing the field in error messaging
type ClubMember struct {
	FirstName      string `spec:"First Name"`
	MiddleName     string
	LastName       string
	CreditScore    int
	Age            int
	MemberSince    time.Time
	MemberExpireAt time.Time
	Guardian       string
}

// ClubMemberSpec defines a specification for a customer
type ClubMemberSpec struct {
	gospecexpress.Specification
}

func newClubMemberSpec() *ClubMemberSpec {
	s := &ClubMemberSpec{}

	s.ForType(&ClubMember{}).
		Custom(func(thing interface{}, contextData map[string]interface{}) (valErr, err error) {
			return fmt.Errorf("This is a test custom validation error returned for the structure as a whole"), nil
		}).
		Required("FirstName").MaxLength(50).
		Optional("MiddleName").MaxLength(20).
		Required("LastName", gospecexpress.WithErrorMessage("Sir Name is a required field!")).MaxLength(50).
		Required("CreditScore").GreaterThan(640, gospecexpress.WithWarning()).
		Required("MemberExpireAt").GreaterThanOtherField("MemberSince").
		Required("Guardian").If(func(t interface{}, c map[string]interface{}) bool { return isMinor(t, c) }).MaxLength(10).
		Optional("Guardian").If(func(t interface{}, c map[string]interface{}) bool { return !isMinor(t, c) }).MaxLength(10)

	return s
}

func isMinor(thing interface{}, contextData map[string]interface{}) bool {
	if cm, ok := thing.(*ClubMember); ok {
		return cm.Age < 18
	}
	return false
}

func init() {
	// Register the ClubMemberSpec in the Catalog.
	// Registering on init ensures the catalog is initialized, however, registration can happen anytime before the catalog is used.
	catalog.Register(newClubMemberSpec())
}

func Benchmark_Valid_ClubMember_Spec(b *testing.B) {
	// Setup
	c := &ClubMember{
		FirstName:      "",
		LastName:       "Flinstone",
		Age:            20,
		CreditScore:    600,
		Guardian:       "A name that is way too long",
		MemberSince:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		MemberExpireAt: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	for i := 0; i < b.N; i++ {
		catalog.Validate(c)
	}
}
