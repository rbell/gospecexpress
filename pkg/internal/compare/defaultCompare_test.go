package compare

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultComparer_Compare_Int32_EqShouldReturnZero(t *testing.T) {
	// setup
	a := int32(123)
	b := int32(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(0), result)
}

func TestDefaultComparer_Compare_Int32_LtShouldReturnNegOne(t *testing.T) {
	// setup
	a := int32(122)
	b := int32(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(-1), result)
}

func TestDefaultComparer_Compare_Int32_GtShouldReturnOne(t *testing.T) {
	// setup
	a := int32(124)
	b := int32(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(1), result)
}

func TestDefaultComparer_Compare_Int64_EqShouldReturnZero(t *testing.T) {
	// setup
	a := int64(123)
	b := int64(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(0), result)
}

func TestDefaultComparer_Compare_Int64_LtShouldReturnNegOne(t *testing.T) {
	// setup
	a := int64(122)
	b := int64(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(-1), result)
}

func TestDefaultComparer_Compare_Int64_GtShouldReturnOne(t *testing.T) {
	// setup
	a := int64(124)
	b := int64(123)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(1), result)
}

func TestDefaultComparer_Compare_String_EqShouldReturnZero(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := "abc"
	//nolint:goconst // testing
	b := "abc"
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(0), result)
}

func TestDefaultComparer_Compare_String_LtShouldReturnNegOne(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := "abc"
	//nolint:goconst // testing
	b := "def"
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(-1), result)
}

func TestDefaultComparer_Compare_String_GtShouldReturnOne(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := "def"
	//nolint:goconst // testing
	b := "abc"
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(1), result)
}

func TestDefaultComparer_Compare_Date_EqShouldReturnOne(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := time.Now()
	//nolint:goconst // testing
	b := a
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(0), result)
}

func TestDefaultComparer_Compare_Date_LTShouldReturnOne(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := time.Now()
	//nolint:goconst // testing
	b := time.Now().Add(time.Hour)
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(-1), result)
}

func TestDefaultComparer_Compare_Date_GtShouldReturnOne(t *testing.T) {
	// setup

	//nolint:goconst // testing
	a := time.Now().Add(time.Hour)
	//nolint:goconst // testing
	b := time.Now()
	aComp := NewDefaultComparer(a)

	// test
	result, err := aComp.Compare(b)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, int(1), result)
}
