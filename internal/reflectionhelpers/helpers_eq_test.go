package reflectionhelpers

import (
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"
)

func TestEq_Bool_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := true
	b := true
	c := !a
	d := !b

	// test
	result1, e1 := Eq(reflect.ValueOf(a), reflect.ValueOf(b))
	result2, e2 := Eq(reflect.ValueOf(c), reflect.ValueOf(d))

	// assert
	assert.Nil(t, e1)
	assert.Nil(t, e2)
	assert.True(t, result1)
	assert.True(t, result2)
}

func TestEq_Bool_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := true
	b := true
	c := !a
	d := !b

	// test
	result1, e1 := Eq(reflect.ValueOf(a), reflect.ValueOf(c))
	result2, e2 := Eq(reflect.ValueOf(d), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e1)
	assert.Nil(t, e2)
	assert.False(t, result1)
	assert.False(t, result2)
}

func TestEq_Int_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := 1
	b := 1

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Int_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := 1
	b := 2

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_UInt_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := uint(1)
	b := uint(1)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_UInt_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := uint(1)
	b := uint(2)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_Float_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := 1.12
	b := 1.12

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Float_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := 1.12
	b := 2.34

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_Complex_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := complex64(1.12)
	b := complex64(1.12)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Complex_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := complex64(1.12)
	b := complex64(2.34)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_String_Equal_ReturnsTrue(t *testing.T) {
	// setup
	//nolint:goconst // ignore suggestion of constant
	a := "test1"
	//nolint:goconst // ignore suggestion of constant
	b := "test1"

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_String_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	//nolint:goconst // ignore suggestion of constant
	a := "test1"
	b := "test2"

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_Time_Equal_ReturnsTrue(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Time_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now.Add(time.Second)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_TimeRef_Equal_ReturnsTrue(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now

	// test
	result, e := Eq(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_TimeRef_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	now := time.Now()
	a := now
	b := now.Add(time.Second)

	// test
	result, e := Eq(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_NilPointersToSameType_ReturnsTrue(t *testing.T) {
	// setup
	var a *string = nil
	var b *string = nil

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_NilPointersToDifferntType_ReturnsFalse(t *testing.T) {
	// setup
	var a *string = nil
	var b *int = nil

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_Struct_Equal_ReturnsTrue(t *testing.T) {
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
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Struct_NotEqual_ReturnsFalse(t *testing.T) {
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
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_Ptr_Equal_ReturnsTrue(t *testing.T) {
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
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Ptr_NotEqual_ReturnsFalse(t *testing.T) {
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
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_StructWithComparer_Equals_ReturnsTrue(t *testing.T) {
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
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Decimal_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(10)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_Decimal_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(20)

	// test
	result, e := Eq(reflect.ValueOf(a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}

func TestEq_DecimalRef_Equal_ReturnsTrue(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(10)

	// test
	result, e := Eq(reflect.ValueOf(&a), reflect.ValueOf(&b))

	// assert
	assert.Nil(t, e)
	assert.True(t, result)
}

func TestEq_DecimalRef_NotEqual_ReturnsFalse(t *testing.T) {
	// setup
	a := decimal.NewFromInt(10)
	b := decimal.NewFromInt(20)

	// test
	result, e := Eq(reflect.ValueOf(&a), reflect.ValueOf(b))

	// assert
	assert.Nil(t, e)
	assert.False(t, result)
}
