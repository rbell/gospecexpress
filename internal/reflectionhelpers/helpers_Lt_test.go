package reflectionhelpers

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"
)

func TestLt_Bool_Equal_ReturnsBadComparisonError(t *testing.T) {
	// setup
	a := true
	b := true
	c := !a
	d := !b

	// test
	result1, e1 := Lt(reflect.ValueOf(a), reflect.ValueOf(b))
	result2, e2 := Lt(reflect.ValueOf(c), reflect.ValueOf(d))

	// assert
	assert.NotNil(t, e1)
	assert.NotNil(t, e2)
	assert.False(t, result1)
	assert.False(t, result2)
}

func TestLt_Bool_NotEqual_ReturnsBadComparisonError(t *testing.T) {
	// setup
	a := true
	b := true
	c := !a
	d := !b

	// test
	result1, e1 := Lt(reflect.ValueOf(a), reflect.ValueOf(b))
	result2, e2 := Lt(reflect.ValueOf(c), reflect.ValueOf(d))

	// assert
	assert.NotNil(t, e1)
	assert.NotNil(t, e2)
	assert.False(t, result1)
	assert.False(t, result2)
}

func TestLt_Int_Equal_ReturnsFalse(t *testing.T) {
	// setup
	a := 1
	b := 1

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_Int_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	a := 1
	b := 2

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_UInt_Equal_ReturnsFalse(t *testing.T) {
	// setup
	a := uint(1)
	b := uint(1)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_UInt_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	a := uint(1)
	b := uint(2)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_Float_Equal_ReturnsFalse(t *testing.T) {
	// setup
	a := 1.12
	b := 1.12

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_Float_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	a := 1.12
	b := 2.34

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_Complex_Equal_ReturnsBadComparisonError(t *testing.T) {
	// setup
	a := complex64(1.12)
	b := complex64(1.12)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_Complex_NotEqual_ReturnsBadComparisonError(t *testing.T) {
	// setup
	a := complex64(1.12)
	b := complex64(2.34)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_String_Equal_ReturnsFalse(t *testing.T) {
	// setup
	//nolint:goconst // ignore suggestion of constant
	a := "test1"
	//nolint:goconst // ignore suggestion of constant
	b := "test1"

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_String_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	//nolint:goconst // ignore suggestion of constant
	a := "test1"
	b := "test2"

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_Time_Equal_ReturnsFalse(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_Time_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now.Add(time.Second)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_TimeRef_Equal_ReturnsFalse(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now

	// test
	result, e := Lt(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_TimeRef_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now.Add(time.Second)

	// test
	result, e := Lt(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_NilPointersToSameType_ReturnsBadComparisonError(t *testing.T) {
	// setup
	var a *string = nil
	var b *string = nil

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_NilPointersToDifferntType_ReturnsBadComparisonError(t *testing.T) {
	// setup
	var a *string = nil
	var b *int = nil

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_Struct_Equal_ReturnsBadComparisonError(t *testing.T) {
	// setup
	type testStruct struct {
		name        string
		pointerTest *string
	}
	a := testStruct{
		name:        "test",
		pointerTest: nil,
	}
	b := testStruct{
		name:        "test",
		pointerTest: nil,
	}

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_Struct_NotEqual_ReturnsBadComparisonError(t *testing.T) {
	// setup
	type testStruct struct {
		name        string
		pointerTest *string
	}
	a := testStruct{
		name:        "test",
		pointerTest: nil,
	}
	v := "testValue"
	b := testStruct{
		name:        "test1",
		pointerTest: &v,
	}

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_Ptr_Equal_ReturnsBadComparisonError(t *testing.T) {
	// setup
	type testStruct struct {
		name        string
		pointerTest *string
	}
	// two pointers referencing same thing
	a := &testStruct{
		name:        "test",
		pointerTest: nil,
	}
	b := a

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_Ptr_NotEqual_ReturnsBadComparisonError(t *testing.T) {
	// setup
	type testStruct struct {
		name        string
		pointerTest *string
	}
	// two pointers referencing different places in memory
	a := &testStruct{
		name:        "test",
		pointerTest: nil,
	}
	b := &testStruct{
		name:        "test",
		pointerTest: nil,
	}

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.NotNil(t, e)
	assert.False(t, result)
}

func TestLt_StructWithComparer_Equals_ReturnsFalse(t *testing.T) {
	// setup
	type testStruct struct {
		name string
		rank int
	}
	a := &testStruct{
		name: "test",
		rank: 1,
	}
	b := &testStruct{
		name: "test",
		rank: 1,
	}
	comp := func(a, b interface{}) int {
		//nolint:errcheck //ignore error
		atest := a.(*testStruct)
		//nolint:errcheck //ignore error
		btest := b.(*testStruct)
		if atest.rank < btest.rank {
			return -1
		}
		if atest.rank > btest.rank {
			return 1
		}
		return 0
	}
	RegisterComparer(reflect.TypeOf(a), comp)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_StructWithComparer_NotEquals_ReturnsTrue(t *testing.T) {
	// setup
	type testStruct struct {
		name string
		rank int
	}
	a := &testStruct{
		name: "test",
		rank: 1,
	}
	b := &testStruct{
		name: "test",
		rank: 2,
	}
	comp := func(a, b interface{}) int {
		//nolint:errcheck //ignore error
		atest := a.(*testStruct)
		//nolint:errcheck //ignore error
		btest := b.(*testStruct)
		if atest.rank < btest.rank {
			return -1
		}
		if atest.rank > btest.rank {
			return 1
		}
		return 0
	}
	RegisterComparer(reflect.TypeOf(a), comp)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_Decimal_Equal_ReturnsFalse(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(10)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_Decimal_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(20)

	// test
	result, e := Lt(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestLt_DecimalRef_Equal_ReturnsFalse(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(10)

	// test
	result, e := Lt(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestLt_DecimalRef_NotEqual_ReturnsTrue(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(20)

	// test
	result, e := Lt(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}
