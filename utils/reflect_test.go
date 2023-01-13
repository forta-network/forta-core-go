package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testInterface interface {
	Foo() int
}

type testImpl1 struct{}

func (Impl1 *testImpl1) Foo() int {
	return 111
}

type testImpl2 struct{}

func (impl2 *testImpl2) Foo() int {
	return 222
}

type testStruct struct {
	Impl1 *testImpl1
	Impl2 *testImpl2
}

func TestGetImplementation(t *testing.T) {
	r := require.New(t)

	s1 := testStruct{
		Impl1: &testImpl1{},
	}
	i, ok := GetImplementation[testInterface](s1)
	r.True(ok)
	r.Equal(111, i.Foo())

	s2 := testStruct{
		Impl1: &testImpl1{},
		Impl2: &testImpl2{},
	}
	i, ok = GetImplementation[testInterface](s2)
	r.True(ok)
	r.Equal(111, i.Foo())

	s3 := testStruct{
		Impl2: &testImpl2{},
	}
	i, ok = GetImplementation[testInterface](s3)
	r.True(ok)
	r.Equal(222, i.Foo())
}
